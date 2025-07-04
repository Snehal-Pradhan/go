package main

import "fmt"

// for -> only construct in go for looping

func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/*infinte loop

	for{

	}
	*/

	// classic loop
	for i := 1; i <= 3; i++ {

		// break and continue also available
		fmt.Println(i)

	}

	// new feature - range
	for i := range 15 {
		fmt.Println(i)
	}

}
