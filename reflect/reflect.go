package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

type MyStruct struct {
	Name string
	Age  int64
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = int64(23)

	result := &MyStruct{}
	err := result.FillStruct(myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println("\nAnd now for some random crap\n")

	str := "I'm a string"
	things := map[int]interface{}{
		1: []string{"a", "b", "c"},
		2: 1.2,
		3: "Something",
		4: &str,
	}

	for _, thing := range things {
		describe(thing)
	}
}

func describe(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()
	k := v.Kind()

	fmt.Println("i = ", v, "of type", t, "of kind", k)
	return
}
