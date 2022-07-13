package main

import (
	"fmt"
)

//Realitza una estructora for que permeti recorrer tot l'array i finalitzi quan ja no hi hagin més elements per recorrer.

func main() {
	alumnes := []string{"Anna", "Pep", "Iria", "Oscar"}
	i := 1
	for {
		switch {
		case i < len(alumnes):
			fmt.Println(alumnes[i])
		default:
			break
		}
	}
}
