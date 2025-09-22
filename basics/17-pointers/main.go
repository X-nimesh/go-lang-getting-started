package main

import "fmt"

func zeroVal(val int)  {
	val=0
}

func zeroPtr(valPtr *int)  {
	*valPtr=0
}

func main() {
  i:=10;
	fmt.Println("Initial:",i);

	zeroVal(i);
	fmt.Println("Zero by Value",i)

	zeroPtr(&i)
	fmt.Println("zero by Ref",i);
}
