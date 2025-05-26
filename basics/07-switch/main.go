package main

import (
	"fmt"
	"time"
)

func main()  {
	i:=2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}


	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend. 🔥")
	default:
		fmt.Println("It's WeekDay.☹️  🥲")
	}
}
