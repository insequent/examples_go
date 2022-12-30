package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Description string
	Fruits      map[string][]string
}

func main() {
	filename, _ := filepath.Abs("./fruits.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Config Marshalled as YAML:")
	y, err := yaml.Marshal(config)
	fmt.Println(string(y))

	fmt.Printf("Value: %#v\n", config)
	fmt.Printf("Description: %v\n", config.Description)
	fmt.Println("Fruits:")
	for key, value := range config.Fruits {
		fmt.Printf("\t%s: %q\n", key, value)
	}
}
