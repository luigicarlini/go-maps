package main

import (
	"fmt"
)

func main() {
	var data = map[string]map[string]string{}
	data["a"] = map[string]string{}
	data["b"] = make(map[string]string)
	data["c"] = make(map[string]string)
	data["a"]["w"] = "x"
	data["b"]["w"] = "y"
	data["c"]["w"] = "z"
	fmt.Println(data)

	var datamap = map[string]map[string]string{
		"a": map[string]string{},
		"b": map[string]string{},
		"c": map[string]string{},
	}
	datamap["a"]["w"] = "alfa"
	datamap["b"]["w"] = "gamma"
	datamap["c"]["w"] = "delta"
	fmt.Println(datamap)

	var datamap2 = map[string]map[string]string{"a": {"w": "charlie"}, "b": {"w": "passo"}, "c": {"w": "chiudo"}}
	fmt.Println(datamap2)

}
