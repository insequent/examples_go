package main

import (
	"encoding/json"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// NOTE: JSON keeps capitalization but YAML does not
type Compaction struct {
	Count          int                `json:"" yaml:""` // When empty, key == struct key
	ErrCount       int                `json:"ErrorCount" yaml:"ErrorCount"`
	ErrPreviousRun bool               `json:"error_previous_run,omitempty" yaml:"error_previous_run,omitempty"` // omitempty removes field from output when empyt
	Running        bool               `json:"running" yaml:"running"`
	Admin          bool               `json:"-" yaml:"-"` // Will be omitted when marshalled
	Tags           map[string]*string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func main() {
	test := []byte(`{"random":{"crap":"yes","needed":false},"compaction":{"count":2,"error_count":3,"running":false}}`)

	var m struct {
		Compaction Compaction
	}

	if err := json.Unmarshal(test, &m); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonM, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	yamlM, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Marshalled JSON:\n%v\n", string(jsonM))
	fmt.Printf("Marshalled YAML:\n%v\n", string(yamlM))
	fmt.Printf("Raw Print:\n%+v\n", m)
}
