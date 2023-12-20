package main

import "fmt"

func main() {
	fmt.Println("Structs is golang")
	//no inheritance in golang, no super, no child or parent
	ashish := User{"Ashish", "ashish@go.dev", true, 22}
	fmt.Println(ashish)
	fmt.Printf("Ashish Details are: %+v\n", ashish)
	ashish.GetStatus()
	ashish.NewMail()
}

//first letter is capital always
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user: ", u.Email)
}
