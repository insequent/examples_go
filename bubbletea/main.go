package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TimeMsg struct{ time time.Time }

type Clock struct {
	done    chan struct{}
	ticker  *time.Ticker
	program *tea.Program
}

func NewClock(p *tea.Program) *Clock {
	clock := &Clock{
		done:    make(chan struct{}),
		program: p,
		ticker:  time.NewTicker(time.Second),
	}
	go clock.Start()

	return clock
}

func (c Clock) Start() {
	for {
		select {
		case t := <-c.ticker.C:
			c.program.Send(TimeMsg{time: t})
		case <-c.done:
			return
		}
	}
}

func (c Clock) Stop() {
	c.done <- struct{}{}
}

type TickMsg struct {
	time string
}

type model struct {
	choices  []string         // items in the todo list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	time     time.Time
}

func newModel() model {
	return model{
		choices:  []string{"Purple", "Red", "Blue", "Green", "Orange", "Yellow"},
		selected: map[int]struct{}{},
		time:     time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	case TimeMsg:
		m.time = msg.time
	}

	return m, nil
}

func (m model) View() string {
	s := "What are your favorite colors?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += fmt.Sprintf("\nClock %s\n", m.time.Format(time.RFC1123))
	s += "\nPress q to quit.\n"

	return s
}

func main() {
	m := newModel()
	p := tea.NewProgram(m)
	clock := NewClock(p)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	clock.Stop()
}
