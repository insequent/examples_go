package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Resource struct {
	Tags []*Tag
}

type Tag struct {
	Key   string
	Value string
}

func GetTagValue(resource interface{}, key string) (value string, err error) {
	r := reflect.ValueOf(resource)
	if r.Kind() == reflect.Ptr {
		r = reflect.Indirect(r)
	}

	if r.Kind() != reflect.Struct {
		err = errors.New("Passed resource is not struct and cannot contain a 'Tags' field")
		return
	}

	tags := r.FieldByName("Tags")

	if tags.Kind() != reflect.Slice {
		err = errors.New("No Tags field found")
		return
	}

	for i := 0; i < tags.Len(); i++ {
		tag := tags.Index(i)
		if tag.Kind() == reflect.Ptr {
			tag = reflect.Indirect(tag)
		}

		k := tag.FieldByName("Key")
		v := tag.FieldByName("Value")
		if k.Kind() != reflect.String || v.Kind() != reflect.String {
			continue
		}

		if strings.ToLower(k.Interface().(string)) == strings.ToLower(key) {
			value = v.Interface().(string)
			return
		}
	}

	return
}

func main() {

	resource := &Resource{
		Tags: []*Tag{
			&Tag{
				Key:   "email",
				Value: "test@test.com",
			},
			&Tag{
				Key:   "name",
				Value: "tester",
			},
		},
	}

	for _, x := range []string{"name", "email", "not real"} {
		val, err := GetTagValue(resource, x)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}
}
