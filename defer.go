package main

import (
	"fmt"
)

func main() {
	// defer keyword defers execution of func until wrapper func returns
	defer fmt.Println("called last")

	fmt.Println("called 1st")
}
