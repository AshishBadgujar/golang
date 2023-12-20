package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string // always have to give size of array

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist is: ", fruitList)
	fmt.Println("Fruitlist is: ", len(fruitList))

}
