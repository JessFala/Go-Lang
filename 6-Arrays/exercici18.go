package main

import (
	"fmt"
)

//Mira de recorrer i mostrar les dades através d'un for amb range.

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}

	for i, v := range carro {
		if i%2 == 0 {
			fmt.Println(i, " ", v)
		}
	}
}
