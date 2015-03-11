package main

import (
	"fmt"
)

func main() {
	// defer keyword defers execution of func until wrapper func returns
	// defer's are pushed onto a stack
	defer fmt.Println("called last")
	defer fmt.Println("called 3rd")
	defer fmt.Println("called 2nd")

	fmt.Println("called 1st")
}
