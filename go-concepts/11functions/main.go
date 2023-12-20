package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	sum, message := proAdder(2, 3, 4, 5, 6)

	fmt.Printf("%s, %v", message, sum)

}

func proAdder(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum, "This is sum"
}
