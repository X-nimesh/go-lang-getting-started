package main

import "fmt"

func multiVal()(int,int){
	return 5,6
}
func main(){
	a,b:=multiVal()
	fmt.Println(a,b)
	_,c:=multiVal()
	fmt.Println(c)
}