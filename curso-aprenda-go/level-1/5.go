// Link https://play.golang.org/p/aaBlJ4zNVoD

package main

import (
	"fmt"
)

type newtype int

var x newtype
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)
	y = int(x)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
}
