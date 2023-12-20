package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If Else in golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	//old way
	// rand.Seed(time.Now().UnixNano())

	//new way
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	diceNumber := randomGenerator.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println(diceNumber)
	case 2:
		fmt.Println(diceNumber)
	case 3:
		fmt.Println(diceNumber)
	case 4:
		fmt.Println(diceNumber)
	case 5:
		fmt.Println(diceNumber)
	case 6:
		fmt.Println(diceNumber)
	default:
		fmt.Println("No choice")
	}
}
