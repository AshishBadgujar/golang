package main

import "fmt"

func main() {
	x := 10
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}

	names := []string{"A", "b", "c", "v"}
	for index, value := range names {
		//value is local here
		fmt.Println(index, value)
	}
	for _, value := range names {
		fmt.Println(value)
	}
}
