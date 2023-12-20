package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"Monday", "Tuesday", "wednesday", "Thursday", "friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index, day)
	// }

	incValue := 1
	for incValue < 10 {
		if incValue == 2 {
			goto lco
		}
		if incValue == 5 {
			break
		}
		fmt.Println("value is: ", incValue)
		incValue++
	}

lco:
	fmt.Println("Jumping at ...")
}
