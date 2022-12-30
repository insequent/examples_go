package main

import (
	"testing"
)

func BenchmarkUnbufferedChanWithEmptyStruct(b *testing.B) {
	ch := make(chan struct{})

	go func() {
		for {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}

func BenchmarkUnbufferedChanWith16ByteSlice(b *testing.B) {
	ch := make(chan []byte)

	go func() {
		for {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- make([]byte, 16)
	}
}

func BenchmarkUnbufferedChanWith64KByteSlice(b *testing.B) {
	ch := make(chan []byte)

	go func() {
		for {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- make([]byte, 64000)
	}
}
