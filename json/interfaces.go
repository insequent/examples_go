package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ResolverConfig interface{}

type Config interface{}

// Actual data type
//type Config struct {
//	Masters []string
//	Timeout int
//	Builtin ResolverConfig
//	Consul  ResolverConfig
//}

func main() {
	var c, d interface{}

	file, err := filepath.Abs("./test.json")

	if err != nil {
		fmt.Printf("Cannot find configuration file")
	} else if bs, err := ioutil.ReadFile(file); err != nil {
		fmt.Printf("Missing configuration file: %q", file)
	} else if err = json.Unmarshal(bs, &c); err != nil {
		fmt.Printf("Failed to unmarshal config file %q: %v", file, err)
	} else if err = json.Unmarshal(bs, &d); err != nil {
		fmt.Printf("Failed to unmarshal config file the second time %q: %v", file, err)
	}

	fmt.Println("JSON output for c: \n", c)
	fmt.Println("JSON output for d: \n", d)

	for k, v := range c.(map[string]interface{}) {
		maps := make(map[string]map[string]interface{})

		switch k {
		case "builtin":
			maps[k] = v.(map[string]interface{})
		case "consul":
			maps[k] = v.(map[string]interface{})
		default:
			fmt.Println(k, v)
		}

		for a, b := range maps {
			fmt.Println(a)
			for x, y := range b {
				fmt.Println("    ", x, y)
			}
		}
	}

}
