package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("INfinite LOOp")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
