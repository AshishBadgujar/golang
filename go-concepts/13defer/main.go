package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myRange()
}

//defer executes at the end of function
//defers executes as last in first out (LIFO)
func myRange() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
