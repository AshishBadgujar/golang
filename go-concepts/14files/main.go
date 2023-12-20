package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This is file content it needs to be in the file."

	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of file: ", length)
	ReadFile("./test.txt")
	defer file.Close()
}

func ReadFile(name string) {
	databyte, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println("File content: ", string(databyte))
}
