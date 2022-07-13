//Amb aquest exercici tens que realitzar una estructura condicional que representi un porter de discoteca que nomès deixa passar si ets major d'edat. Pero ara amb la possibilitat de donar una resposta per quan no pot passar.

package main

import (
	"fmt"
)

func main() {
	edat := 19
	if edat >= 18 {
		fmt.Println("Pots accedir") //Missatge per quan es compleix la condició
	} else {
		fmt.Println("No pots passar") //Missatge per quan no es compleix la condició
	}
}
