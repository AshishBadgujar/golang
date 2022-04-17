package main

import "fmt"

func main() {
	age := 78
	if age < 30 {
		fmt.Println("less than 30")
	} else if age < 40 {
		fmt.Println("lesss than 40")
	} else {
		fmt.Println("Greater than 30")
	}

	names := []string{"mario", "ashish", "bharagv", "ninja"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("value at pos %v is %v \n", index, value)
	}
}
