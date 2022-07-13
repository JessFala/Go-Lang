package Main

import (
	"fmt"
)

type persona struct { //STRUC persona amb dos camps: un de tipus nom (string) i un de tipus edat (int)
	nom  string
	edat int
}

type ingenier struct { //Per NO repetir cridem l'struc persona dins l'struc ingenier i li hem afegit un boolean.
	persona
	realitzarPlanol bool
}

func main() {
	ing1 := ingenier{
		persona: persona{
			nom:  "Ruben",
			edat: 26,
		},
		realitzarPlanol: true, //coma (,) també al final!
	}

	fmt.Println(ing1)
	fmt.Println(ing1.nom, ing1.edat, ing1.realitzarPlanol) //Els atributs de persona i d'edat són atributs que s'invoquen des d'ing1; perquè en realitat provenen 'persona' i no d''ing1')
}
