package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")
	presentTime := time.Now()
	fmt.Println("present Time:", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// 01 month
	// 02 day
	// 2006 year
	// Monday = day

	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
