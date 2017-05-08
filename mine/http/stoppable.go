// This is stolen from https://github.com/facebookgo/httpdown/blob/master/httpdown.go

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type Server interface {
    Wait() error
    Stop() error
}


type HTTP struct {
    stopTimeout time.Duration
    killTimeout time.Duration
}

func (h HTTP) Serve(s *http.Server, l net.Listener) Server {
    ss := &server{
        active:       make(chan net.Conn),
        closed:       make(chan net.Conn),
        idle:         make(chan net.Conn),
        kill:         make(chan chan struct{}),
        killTimeout:  h.killTimeout,
        listener:     l,
        new:          make(chan net.Conn),
        oldConnState: s.ConnState,
        stopTimeout:  h.stopTimeout,
        serveDone:    make(chan struct{}),
        serveErr:     make(chan error, 1),
        server:       s,
        stop:         make(chan chan struct{}),
    }

    s.ConnState = ss.connState
    go ss.manage()
    go ss.serve()
    return ss
}

func (h HTTP) ListenAndServe(s *http.Server) (Server, error) {
	addr := s.Addr
	if addr == "" {
		if s.TLSConfig == nil {
			addr = ":80"
		} else {
			addr = ":443"
		}
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	if s.TLSConfig != nil {
		l = tls.NewListener(l, s.TLSConfig)
	}
	return h.Serve(s, l), nil
}

type server struct {
    active       chan net.Conn
    closed       chan net.Conn
    idle         chan net.Conn
    kill         chan chan struct{}
    killTimeout  time.Duration
    listener     net.Listener
    new          chan net.Conn
    oldConnState func(net.Conn, http.ConnState)
    serveDone    chan struct{}
    serveErr     chan error
    server       *http.Server
    stopTimeout  time.Duration
    stop         chan chan struct{}
    stopErr      error
    stopOnce     sync.Once
}

func (s *server) connState(c net.Conn, cs http.ConnState) {
	if s.oldConnState != nil {
		s.oldConnState(c, cs)
	}

	switch cs {
	case http.StateNew:
		s.new <- c
	case http.StateActive:
		s.active <- c
	case http.StateIdle:
		s.idle <- c
	case http.StateHijacked, http.StateClosed:
		s.closed <- c
	}
}

func (s *server) manage() {
	defer func() {
		close(s.new)
		close(s.active)
		close(s.idle)
		close(s.closed)
		close(s.stop)
		close(s.kill)
	}()

	var stopDone chan struct{}

	conns := map[net.Conn]http.ConnState{}
	var countNew, countActive, countIdle float64

	decConn := func(c net.Conn) {
		switch conns[c] {
		default:
			panic(fmt.Errorf("unknown existing connection: %s", c))
		case http.StateNew:
			countNew--
		case http.StateActive:
			countActive--
		case http.StateIdle:
			countIdle--
		}
	}

	for {
		select {
		case c := <-s.new:
			conns[c] = http.StateNew
			countNew++
		case c := <-s.active:
			decConn(c)
			countActive++

			conns[c] = http.StateActive
		case c := <-s.idle:
			decConn(c)
			countIdle++

			conns[c] = http.StateIdle

			if stopDone != nil {
				c.Close()
			}
		case c := <-s.closed:
			decConn(c)
			delete(conns, c)

			if stopDone != nil && len(conns) == 0 {
				close(stopDone)
				return
			}
		case stopDone = <-s.stop:
			if len(conns) == 0 {
				close(stopDone)
				return
			}

			for c, cs := range conns {
				if cs == http.StateIdle {
					c.Close()
				}
			}

		case killDone := <-s.kill:
			for c := range conns {
				c.Close()
			}

			close(killDone)
		}
	}
}

func (s *server) serve() {
	s.serveErr <- s.server.Serve(s.listener)
	close(s.serveDone)
	close(s.serveErr)
}

func (s *server) Wait() error {
	if err := <-s.serveErr; !isUseOfClosedError(err) {
		return err
	}
	return nil
}

func (s *server) Stop() error {
	s.stopOnce.Do(func() {
		s.server.SetKeepAlivesEnabled(false)

		closeErr := s.listener.Close()
		<-s.serveDone

		stopDone := make(chan struct{})
		s.stop <- stopDone

		if closeErr != nil && !isUseOfClosedError(closeErr) {
			s.stopErr = closeErr
		}
	})
	return s.stopErr
}

func isUseOfClosedError(err error) bool {
	if err == nil {
		return false
	}
	if opErr, ok := err.(*net.OpError); ok {
		err = opErr.Err
	}
	return err.Error() == "use of closed network connection"
}

func ListenAndServe(s *http.Server, hd *HTTP) error {
	if hd == nil {
		hd = &HTTP{}
	}
	hs, err := hd.ListenAndServe(s)
	if err != nil {
		return err
	}

	waiterr := make(chan error, 1)
	go func() {
		defer close(waiterr)
		waiterr <- hs.Wait()
	}()

	signals := make(chan os.Signal, 10)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-waiterr:
		if err != nil {
			return err
		}
	case <-signals:
		signal.Stop(signals)
		if err := hs.Stop(); err != nil {
			return err
		}
		if err := <-waiterr; err != nil {
			return err
		}
	}
	return nil
}





