package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:",a)

	a[4]=100
	fmt.Println("set:",a)

	b:=[5]int{1,2,3,4}
	fmt.Println("dcl:",b)
	twoD:= [2][3]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(twoD)
}