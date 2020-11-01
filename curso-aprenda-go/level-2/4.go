/**
  - Atribua um valor int a uma variável
  - Demonstre este valor em decimal, binário e hexadecimal
  - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
  - Demonstre esta outra variável em decimal, binário e hexadecimal

  Link https://play.golang.org/p/-DnJN-qoRzT
*/

package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Printf("%v - %b - %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%v - %b - %#x", y, y, y)
}
