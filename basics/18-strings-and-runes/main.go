package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
const n="नमस्कार";
	fmt.Println(len(n))

	for i:=0;i<len(n);i++{
			fmt.Printf("%x",n[i])
	}
	fmt.Println("Rune count:",utf8.RuneCountInString(n))
	
	for idx, runeValue:= range n{
		fmt.Printf("%#U starts at %d\n",runeValue,idx)
	}

}
