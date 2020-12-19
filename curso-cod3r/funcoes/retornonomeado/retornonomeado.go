package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	segundo, primeiro := trocar(1, 2)
	fmt.Println(segundo, primeiro)
}
