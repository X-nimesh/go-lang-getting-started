package main

import (
	"fmt"
)

func main() {
	nums:=[]int {2,3,4}
	sum:=0
	for index,num:= range nums {
		fmt.Println(index,num)
		sum+=num;
	}

	strArr:=map[string]string{"nam":"nimesh","age":"23"}
	for index,value:=range strArr{
		fmt.Println(index,":",value)
	}
	fmt.Println("Sum:",sum);
}
