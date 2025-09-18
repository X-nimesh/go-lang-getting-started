package main

import "fmt"


func closureInt()func()int  {
	i:=0
	return func () int {
		i++
		return i
	}
	
}
func main() {
	nextInt:=closureInt()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
