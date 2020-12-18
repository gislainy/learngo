package main

import "fmt"

func main() {
	funcsESSalarios := map[string]float64{
		"João":  1135.45,
		"Maria": 1556.78,
		"Pedro": 1200.0,
	}

	funcsESSalarios["Rafael"] = 1350.0
	delete(funcsESSalarios, "Inexistente")

	for nome, salario := range funcsESSalarios {
		fmt.Println(nome, salario)
	}

	fmt.Println(funcsESSalarios)
}
