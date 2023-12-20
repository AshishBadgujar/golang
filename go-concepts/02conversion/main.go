package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion.")
	fmt.Println("Rate my pizza out of 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Rating type is %T", input)

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("something went wrong")
	} else {
		fmt.Println("Added 1 to your rating: ", numberRating+1)
	}

}
