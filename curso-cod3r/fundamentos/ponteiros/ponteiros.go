package main

import "fmt"

func main() {
	i := 1

	// Go nã otem aritmética de ponteiros

	var p *int = nil

	p = &i //pegando o endereço da variavel

	*p++ //desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, i, &i)

}
