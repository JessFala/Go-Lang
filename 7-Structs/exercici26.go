/*
TODO A continuació realitza 2 structs anònims sobre articles de supermercats a on
TODO has de definir les dades de nom, unitats i pvp.*/

package main

import (
	"fmt"
)

func main() {
	article1 := struct {
		nom          string
		unitats, pvp float32
	}{
		nom:     "Vi negre",
		unitats: 3,
		pvp:     5.50,
	}

	article2 := struct {
		nom          string
		unitats, pvp float32
	}{
		nom:     "Formatge",
		unitats: 2,
		pvp:     6.3,
	}

	article3 := struct {
		nom          string
		unitats, pvp float32
	}{
		nom:     "Olives",
		unitats: 1,
		pvp:     4.5,
	}

	fmt.Println(article1)
	fmt.Println(article2)
	fmt.Println(article3)
}
