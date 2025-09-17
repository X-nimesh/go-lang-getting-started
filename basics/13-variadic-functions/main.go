package main

import "fmt"


func sumMultiple(nums ...int)  {
	fmt.Println(nums);
	total:=0;
	for _,num:=range nums{
		total+=num;
	}
	fmt.Println("Total Sum:",total)
}


func main()  {
		sumMultiple(1,2,2)	
} 
