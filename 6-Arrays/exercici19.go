package main

import (
	"fmt"
)

/*Tenint en compte que la col.lecció de dades que mostrem,
correspon a la facturació d'una tenda de cada mes del any,
tens que mostrar el total facturat el primer trimestre.*/

func main() {
	carro := []int{23452, 72622, 5433, 9273, 62533, 62353, 15625, 25162, 41312, 62762, 62532, 52426}
	total := 0
	fmt.Println(carro[0:3])
	for _, v := range carro {
		total += v
	}

	fmt.Println("El total del primer trimestre és ", total)
}
