package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("o valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	bo := true

	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	s1 := "Olá meu nome é Gi"
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Ola
	meu
	nome
	é
	Gi`
	fmt.Println("O tamanho da string é", len(s2))

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

}
