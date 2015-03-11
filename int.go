package main

import (
	"fmt"
)

/**
 * byte is alias for uint8
 */
var uintVar uint8 = 1
var byteVar byte = uintVar

/**
 * rune is alias for int32
 */
var int32Var int32 = 41
var runeVar rune = int32Var

func main() {
	fmt.Println(byteVar)
	fmt.Println(runeVar)
}
