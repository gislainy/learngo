/**
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
Link https://play.golang.org/p/I4TicVmYUYF
*/

package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Printf("%v - %b - %#x", x, x, x)
}
