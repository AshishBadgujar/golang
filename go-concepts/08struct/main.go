package main

import "fmt"

func main() {
	fmt.Println("Structs is golang")
	//no inheritance in golang, no super, no child or parent
	ashish := User{"Ashish", "ashish@go.dev", true, 22}
	fmt.Println(ashish)
	fmt.Printf("Ashish Details are: %+v\n", ashish)
}

//first letter is capital always
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
