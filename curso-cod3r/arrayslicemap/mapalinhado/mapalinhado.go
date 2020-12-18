package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga":          8446.78,
		},
		"J": {
			"Jose": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {

		fmt.Printf("\nLetra %v: ", letra)

		for nome, salario := range funcs {
			fmt.Printf("\n%s - %.2f", nome, salario)
		}
	}

}
