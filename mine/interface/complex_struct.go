package main

import (
	"fmt"
)

func NewWatcher(r Resource) *Watcher {
	fmt.Println("CREATE: ", r.GetName())
	return &Watcher{r: r}
}

type Watcher struct {
	r Resource
}

func (w *Watcher) GetName() string {
	return w.r.GetName()
}

type Resource interface {
	Add() error
	GetName() string
	Remove() error
}

type Thing1 struct {
	Name string
}

func (t *Thing1) Add() error {
	fmt.Println("ADD: ", t.Name)
	return nil
}

func (t *Thing1) GetName() string {
	return t.Name
}

func (t *Thing1) Remove() error {
	fmt.Println("REMOVE: ", t.Name)
	return nil
}

type Thing2 struct {
	Name string
}

func (t *Thing2) Add() error {
	fmt.Println("ADD: ", t.Name)
	return nil
}

func (t *Thing2) GetName() string {
	return t.Name
}

func (t *Thing2) Remove() error {
	fmt.Println("REMOVE: ", t.Name)
	return nil
}

func main() {
	var watcher *Watcher
	thing1 := &Thing1{Name: "thing1"}
	thing2 := &Thing2{Name: "thing2"}

	for _, resource := range []Resource{thing1, thing2} {
		watcher = NewWatcher(resource)
		fmt.Println("Watcher has a name of", watcher.GetName())
	}
}
