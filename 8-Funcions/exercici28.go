package Main

import (
	"fmt"
)

type alumne struct {
	nom    string
	cognom string
}
type matriculat struct {
	alumne
	asignatura string
}

//Exemple de mètode
func (a alumne) presentarse() { //D’aquesta manera afegim aquest mètode a l'struct matriculat
	fmt.Println("El meu nom és ", a.nom, a.cognom, " i estic aprenent molt.")
}

//D’aquesta manera afegim aquest mètode a l'struct matriculat. És més complet.
func (m matriculat) presentarse() {
	fmt.Println("El meu nom és ", m.nom, m.cognom, " i estic impartint ", m.asignatura)
}

// la Interface obliga a l'ús d'una funció
type usuari interface {
	presentarse()
}

//Ús d'aquesta Interface
func aula(u usuari) {
	fmt.Println("Estic dins de l'aula", u)
}

//Funció ppal amb informació completada.
func Main() {
	al1 := matriculat{
		alumne: alumne{
			nom:    "Jon",
			cognom: "Ruiz",
		},
		asignatura: "Go",
	}

	pr1 := alumne{
		nom:    "Joan",
		cognom: "Riera",
	}

	fmt.Println(al1)  //{{Jon Ruiz} Go}
	al1.presentarse() //El meu nom és Jon Ruiz i estic aprenent Go
	al1.presentarse() //El meu nom és  Joan Riera  i estic aprenent molt.

	aula(pr1) //Estic dins de l'aula {Joan Riera}
	aula(al1) //Estic dins de l'aula {{Jon Ruiz} Go}
}
