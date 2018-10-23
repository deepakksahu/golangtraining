package main

import (
	"fmt"
)

func countryCurrency() {

	m := make(map[string]string)
	m["sgp"] = "sgd"
	m["us"] = "usd"

	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
	x, ok := m["india"]
	fmt.Println(x, ok)

}

func main() {
	//fmt.Println("Hello, playground")
	countryCurrency()
}

