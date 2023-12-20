package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int // pointer which is responsible for holding integer in that
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pounter is ", ptr)
	fmt.Println("Value of actual pounter is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
