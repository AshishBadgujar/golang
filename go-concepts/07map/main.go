package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("Languages: ", languages)
	fmt.Println("JS short for: ", languages["js"])
	delete(languages, "rb")
	fmt.Println("Languages: ", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
