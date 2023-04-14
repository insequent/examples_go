package main

import "fmt"

func updateMap(m map[string]string) {
	for k, v := range m {
		m[k] = "Updated:" + v
	}
}

func main() {
	map1 := make(map[string]string)
	map1["this"] = "test"
	map1["should"] = "work"
	fmt.Println("Map 1:", map1)

	// nil map
	var map2 map[string]interface{}
	if map2 == nil {
		fmt.Println("Map 2 is nil!")
	}
	fmt.Println("Map 2:", map2)

	map3 := map[string]interface{}{
		"this": 3,
		"that": "yep",
	}
	map3["moar"] = "more"
	fmt.Println("Map 3:", map3)

	// Updating map by reference in function
	map4 := map[string]string{
		"A": "Apple",
		"B": "Banana",
	}
	updateMap(map4)
	fmt.Println("Map 4:", map4)
}
