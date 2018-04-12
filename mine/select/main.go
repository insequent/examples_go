package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Baby struct {
	name string

	say    chan string
	listen chan string

	bedtime chan struct{}
}

func (b *Baby) SayStuff() {
	defer close(b.say)

	b.say <- "I can count!"
	for i := 1; i <= 10; i++ {
		b.say <- strconv.Itoa(i)
	}
	b.say <- "DONE!"

	<-b.bedtime
}

func (b *Baby) HearStuff() {
	for {
		select {
		case <-b.bedtime:
			return
		case words := <-b.listen:
			fmt.Printf("%v: I heard: %v\n", b.name, words)
		}
	}
}

func (b *Baby) Tired() {
	fmt.Printf("%v: I'm schleepy...\n", b.name)
}

func NewMama(baby *Baby, giveup chan *Baby) (m *Mama) {
	secondBaby := &Baby{
		name:    "Baby2",
		listen:  make(chan string),
		bedtime: baby.bedtime,
	}

	m = &Mama{
		firstBaby:  baby,
		secondBaby: secondBaby,
		giveup:     giveup,
		bedtime:    baby.bedtime,
	}

	return
}

type Mama struct {
	firstBaby  *Baby
	secondBaby *Baby

	giveup  chan *Baby
	bedtime chan struct{}
}

func (m *Mama) Loop() {
	for {
		select {
		case words := <-m.firstBaby.say:
			m.secondBaby.listen <- words
			if strings.Contains(words, "DONE") {
				m.GiveUp()
				return
			}
		}
	}
}

func (m *Mama) GiveUp() {
	fmt.Println("Mama: I'm DONE!")
	m.giveup <- m.secondBaby
	return
}

type Papa struct {
	baby *Baby
}

func main() {
	t0 := time.Now()

	say := make(chan string)
	giveup := make(chan *Baby)
	bedtime := make(chan struct{})

	papa := &Papa{
		baby: &Baby{
			name:    "Baby1",
			say:     say,
			bedtime: bedtime,
		},
	}

	mama := NewMama(papa.baby, giveup)
	go papa.baby.SayStuff()
	go mama.secondBaby.HearStuff()
	go mama.Loop()

	papa.baby = <-giveup

	papa.baby.Tired()
	close(bedtime)

	fmt.Printf("ALL ASLEEP! Babies were up for %v\n", time.Since(t0))
}
