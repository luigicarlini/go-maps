package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; map literal
	x := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}

	fmt.Printf("%T\n", x) //print Type
	fmt.Println(x["Nina"])

	for k, v := range x {
		fmt.Println(k, "-", v)
	}

	for k, _ := range x {
		fmt.Println(k, "-", x[k])
	}
}
