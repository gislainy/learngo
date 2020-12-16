package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.141
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Print(area)
}
