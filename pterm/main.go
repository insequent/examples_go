package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	var someList = strings.Split("this that paddywack", " ")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(someList)).WithTitle("Downloading stuff...").Start()

	for i := 0; i < p.Total; i++ {
		p.Title = "Downloading " + someList[i]
		pterm.Success.Println("Downloading " + someList[i])
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}
}
