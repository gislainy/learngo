package main

import "fmt"

func inc1(n int) {
	n++
}

//não é mais um função pura
func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n) // por valor, n não sera afeatado fora do contexto
	fmt.Println(n)
	inc2(&n) //por referência
	fmt.Println(n)
}
