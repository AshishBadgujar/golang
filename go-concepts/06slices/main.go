package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Banana", "pineApple"}
	fmt.Printf("Type of fruitlist %T", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 034
	highScores[1] = 934
	highScores[2] = 234
	highScores[3] = 334
	// highScores[4] = 777 error

	highScores = append(highScores, 555, 666, 888)

	fmt.Println(sort.IntsAreSorted(highScores)) //boolean
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //boolean

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
