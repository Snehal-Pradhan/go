package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("other")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("It's workday")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case bool:
			fmt.Println("It's a boolean")
		case string:
			fmt.Println("It's a string")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang")

}
