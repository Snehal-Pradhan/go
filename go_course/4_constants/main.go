package main

import "fmt"

const age = 30
const (
	port = 5000
	host = "localhost"
)

//shorthand dont work here

func main() {
	const name = "golang"
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(port)
	fmt.Println(host)
}
