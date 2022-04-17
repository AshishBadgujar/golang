package main

import "fmt"

func main() {
	// strings
	var nameOne string = "ashish"
	var nameTwo = "badgujar"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)
	nameOne = "bhargav"
	nameThree = "prakash"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameour := "without var keyword"

	//cannot use short hand outside of the function :=
	//cannot use '' in go
	fmt.Println(nameour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory -128 to 127
	var numOne int8 = 34
	var numTwo int8 = -128
	var numThree uint = 25 //cant have -ve number
	fmt.Println(numOne, numTwo, numThree)

	//float
	var floatOne float32 = -32.34
	var floatTwo float64 = 5448849584394532.34
	floatThree := 1.5 //by default float64

	fmt.Printf("first= %v ,second= %v ,thrid= %v \n", floatOne, floatTwo, floatThree)
	// %q,%T

	// Sprintf
	var str = fmt.Sprintf("saved string")
	fmt.Println(str)

	// arrays(fixed lenght)
	// var ages[3]int=[3]int{20,25,30}
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))
	names := [4]string{"Ashish", "ashu", "bhargav", "vishu"}
	names[1] = "badgujar"
	fmt.Println(names, len(names))

	//slices(use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 34
	scores = append(scores, 89)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangethree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangethree)

}
