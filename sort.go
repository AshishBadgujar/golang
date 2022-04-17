package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{2, 3, 4, 5, 6, 76, 7, 8, 7, 43}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 3)
	fmt.Println(index)

	names := []string{"ashi", "mario", "ashu", "vishu"}
	fmt.Println(names)
	strindex := sort.SearchStrings(names, "ashu")
	fmt.Println(strindex)

}
