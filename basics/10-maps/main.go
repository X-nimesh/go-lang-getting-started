package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println("Map1:", m)
	fmt.Println("Value for k1:", m["k1"])
	fmt.Println("Value for k2:", m["k2"])

	fmt.Println("Length:", len(m))

	delete(m, "k2")
	fmt.Println("Map2:", m)

	clear(m)
	fmt.Println(m)
	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
	n1 := map[string]int{"foo": 1, "barr": 2}
	fmt.Println(n1)
	fmt.Println("n==n1:", maps.Equal(n, n1))

	if maps.Equal(n, n1) {
		fmt.Println("n==n1")
	}
}
