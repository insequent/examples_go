package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Compaction struct {
	Count          int                `json:""` // When empty, json output key == struct key
	ErrCount       int                `json:"error_count"`
	ErrPreviousRun bool               `json:"error_previous_run,omitempty"` // omitempty removes field from output when empyt
	Running        bool               `json:"running"`
	Admin          bool               `json:"-"` // Will be omitted when marshalled
	Tags           map[string]*string `json:"tags,omitempty"`
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

	output, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Premarshalled:\n%v\n", string(test))

	fmt.Printf("Marshalled:\n%v\n", string(output))

	fmt.Printf("Raw Print:\n%+v\n", m)

}
