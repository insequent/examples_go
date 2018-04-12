package main

import "fmt"

func main() {
	blah := map[string]interface{}{
		"this": 3,
		"that": "yep",
	}

	blah["moar"] = "more"

	// nil map
	var map2 map[string]interface{}

	fmt.Println(blah)

	if map2 == nil {
		map2 = map[string]interface{}{
			"a": 1,
			"b": true,
		}
	} else {
		map2["a"] = 1
		map2["b"] = true
	}

	fmt.Println(map2)
}
