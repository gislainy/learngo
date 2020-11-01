/**
Crie uma variável de tipo string utilizando uma raw string literal.
Link https://play.golang.org/p/Ulal4oUCXcE
*/

package main

import (
	"fmt"
)

func main() {
	s := `raw 
		string 
			literal`
	fmt.Println(s)
}
