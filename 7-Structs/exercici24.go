package main

import (
	"fmt"
)

type Persona struct { // 'persona' es considerarà un Tipus personalitzat.
	nom string /* Quan es crea una nova entitat estem obligats a utilitzar
	aquestes dues entitats amb aquest tipus de dades (string i int).*/
	edat int
}

func Main() {
	p1 := Persona{
		nom:  "Gerard",
		edat: 22,
	}

	p2 := Persona{
		nom:  "Ona",
		edat: 18, //No es pot deixar buit!!! Donaria error.
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.nom) //Podem accedir només a una propietat d'aquest element.
	fmt.Println(p2.nom)
}
