package main

	
import (
    "fmt"
    // "slices"
)

func main() {
	// var a[5] int
	
	s:= make([]int,3,10)
	s1:= make([]int,1,1)

	fmt.Println("length:", len(s))
	fmt.Println("Capacity",cap(s))

	s2 := [2]int{1,2}
	s2[0]=0
	fmt.Println(s2)

	s3 := append(s, s1...)
	fmt.Println(s3)


 	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}