package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	// returns ptr of reflect.Value
	structValue := reflect.ValueOf(obj)
	// returns ptr of reflect.Value
	structFieldValue := reflect.Indirect(structValue).FieldByName(name)

	// Check if field exists
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	// Check if field is settable
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type did not match struct field type")
	}

	structFieldValue.Set(val)
	return nil
}

func MapToStruct(m map[string]string, obj interface{}) {
	for k, v := range m {
		// Simple example as v is always a string here...
		err := SetField(obj, k, v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	s := struct {
		Blah  string
		Stuff string
		Not   string
	}{"foo", "2", "it"}
	m := map[string]string{"Blah": "blargh", "Stuff": "42", "Extra": "blah"}

	s.Blah = "arrrrggg"

	fmt.Println("Beginning struct:", s)

	MapToStruct(m, &s)

	fmt.Println("Ending struct:", s)
}
