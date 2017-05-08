package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Test struct {
	Start string
}

// value receiver
func (t Test) Finish() string {
	return t.Start + "finish"
}

// pointer receiver
func (t *Test) New() *Test {
	return &Test{}
}

func CallMethod(i interface{}, methodName string) interface{} {
	var finalMethod reflect.Value
	var ptr reflect.Value

	value := reflect.ValueOf(i)

	// We need a value and a ptr for testing the Method type later
	if value.Kind() == reflect.Ptr {
		ptr = value
	} else {
		ptr = reflect.New(reflect.TypeOf(value))
	}
	value = reflect.Indirect(value)

	// Test for Method by value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// Test for Method by reference
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if finalMethod.IsValid() {
		return finalMethod.Call([]reflect.Value{})[0].Interface()
	}

	// return or panic, method not found of either type
	return errors.New("finalMethod.IsValid() == false")
}

func main() {
	i := Test{Start: "start"}
	j := Test{Start: "start2"}

	fmt.Println("i", CallMethod(i, "Finish"))
	fmt.Println("&i", CallMethod(&i, "Finish"))
	fmt.Println("j", CallMethod(j, "Finish"))
	fmt.Println("&j", CallMethod(&j, "Finish"))

	i2 := CallMethod(&i, "New")
	fmt.Printf("&i, i2: %#v\n", i2)
	j2 := CallMethod(&j, "New")
	fmt.Printf("&j, j2: %#v\n", j2)
	i2 = CallMethod(&i, "New")
	fmt.Printf("i, i2: %#v\n", i2)
	j2 = CallMethod(&j, "New")
	fmt.Printf("j, j2: %#v\n", j2)

	j_real := j2.(*Test)
	j3 := j_real.New()
	fmt.Printf("j_real.New(), j3: %#v\n", j3)
}
