package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  2.33,
		"pie":   3.14,
		"salad": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//loooping
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		2342342342:   "ashish",
		342342423423: "vishu",
		9075752323:   "ashu",
	}
	fmt.Println(phonebook)
	phonebook[9075752323] = "badgujar"
	fmt.Println(phonebook)

}
