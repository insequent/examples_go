package main

import (
	"encoding/json"
	"fmt"
)

type Compaction struct {
	Count          int  `json:"count"`
	ErrCount       int  `json:"error_count"`
	ErrPreviousRun bool `json:"error_previous_run"`
	Running        bool `json:"running"`
}

func main() {
	test := []byte(`{"random":{"crap":"yes","needed":false},"compaction":{"count":2,"error_count":3,"running":false}}`)

	var metadata struct {
		Compaction Compaction
	}

	if err := json.Unmarshal(test, &metadata); err != nil {
		fmt.Println(err)
	}

	c := metadata.Compaction

	fmt.Println(c)
	fmt.Println("Count:", c.Count)
	fmt.Println("ErrCount:", c.ErrCount)
	fmt.Println("ErrPreviousRun:", c.ErrPreviousRun)
	fmt.Println("Running:", c.Running)
}
