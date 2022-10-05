package main

import (
	"fmt"
)

func main() {
	// composite literal; map literal
	x := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	for k, v := range x {
		fmt.Println(k, "-", v)
	}

	y := make(map[string]string)
	y["James"] = "007"
	fmt.Println(y["James"])
}
