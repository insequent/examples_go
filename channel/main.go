package main

import "fmt"

type Speaker struct {
	doneChan chan struct{}
	talkChan chan []byte
}

func (s *Speaker) Speak() {
	s.talkChan <- []byte("So one day I was walking down the street.")
	s.talkChan <- []byte("When I heard a loud *bang* behind me!")
	s.talkChan <- []byte("\"OMG! Aliens!\", I exclaim.")
	s.talkChan <- []byte("Alas, it was just a opossum getting dinner in my neighbor's trashcan.")
	s.talkChan <- nil
	s.doneChan <- struct{}{}
}

type Listener struct {
	doneChan chan struct{}
	talkChan chan []byte
}

func (l *Listener) repeat(words []byte) {
	fmt.Printf("I heard: %s\n", words)
}

func (l *Listener) Listen() {
	for words := range l.talkChan {
		if words == nil {
			l.doneChan <- struct{}{}
			return
		}
		l.repeat(words)
	}
}

func main() {
	talkChan := make(chan []byte)
	doneChan := make(chan struct{})

	speaker := &Speaker{
		doneChan: doneChan,
		talkChan: talkChan,
	}
	listener := &Listener{
		doneChan: doneChan,
		talkChan: talkChan,
	}

	go listener.Listen()
	go speaker.Speak()

	<-doneChan
	<-doneChan
	fmt.Println("Main is all done!")
}
