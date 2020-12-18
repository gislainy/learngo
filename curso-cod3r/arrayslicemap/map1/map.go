package main

import "fmt"

func main() {
	//mapas devem ser inicialziados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"

	fmt.Println(aprovados)

	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(aprovados[123456789])
	delete(aprovados, 123456789)
	fmt.Println(aprovados[123456789])

}
