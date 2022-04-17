package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	//the original value is unchanged
	fmt.Println("original=", greeting)
}
