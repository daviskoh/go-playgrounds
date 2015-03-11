package main

import (
	"fmt"
)

func main() {
	// & in front of var means retrieve memory address (1)
	// * in front of type signals declared var will store memory address (2)
	// * in front of variable of pointer type used to retrieve value at memory address (3)

	ten := 10
	var memoryAddress *int = /* (2) */ &ten // (1)

	fmt.Println(memoryAddress)

	fmt.Println(*memoryAddress)

	// mutates ten
	*memoryAddress += 5
	fmt.Println(ten)
}
