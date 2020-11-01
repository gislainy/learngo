/**
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Link https://play.golang.org/p/iE2XIuqyqls
*/
package main

import (
	"fmt"
)

const (
	_    = (2020 + iota)
	ano1 
	ano2
	ano3
	ano4
)

func main() {
	fmt.Println(ano1, ano2, ano3, ano4)
}
