package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("GoodBye %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func getInitials(n string) (string, string) {
	//multiple return
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	sayGreeting("ashish")
	sayBye("Bhargav")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	// multiple return
	fn, sn := getInitials("Ashish badgujar")
	fmt.Println(fn, sn)
}
