package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i <= 10 { //não tem while em Goi
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}

}
