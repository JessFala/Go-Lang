package Main

import (
	"fmt"
)

type alumne struct {
	nom    string
	cognom string
}
type matriculat struct {
	alumne     //Estem cridant al primer struct i per tant invoquem tant el nom com el cognom.
	asignatura string
}

//Exemple de mètode
func (m matriculat) presentarse() { //D’aquesta manera afegim aquest mètode a l'struct matriculat.
	fmt.Println("El meu nom és ", m.nom, m.cognom, " i estic aprenent ", m.asignatura)
}
func main() {
	al1 := matriculat{
		alumne: alumne{
			nom:    "Jon",
			cognom: "Ruiz",
		},
		asignatura: "Go",
	}
	fmt.Println(al1)  //{{Jon Ruiz} Go}
	al1.presentarse() //El meu nom és Jon Ruiz i estic aprenent Go
}
