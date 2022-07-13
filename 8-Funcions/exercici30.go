// Ex. Call backs

package main

import (
	"fmt"
)

func multi(num int) int { //Funció  multiplicar: Jo rebo un número i retorno un número.
	return num * num
}

func cuadrat(f func(int) int, list []int) []int {
	//Cal recordar que Make serveix per crear Slices i s’ha de indicar make(tipus[], longitud, capacitat)
	var a = make([]int, len(list), cap(list)) //len= Longitud; cap= Capacitat
	for index, val := range list {            //Recórrer la llista que estic rebent
		a[index] = f(val) // Igualo l'índex perquè em torni el resultat de la invocació de la funció.
	}
	return a //Autoritzo que retorni una array d'integers
}

func main() {
	llistat := []int{3, 5, 8, 10}
	resultat := cuadrat(multi, llistat) //Retornarà el doble dels números del 'llistat' de la línia 23.
	fmt.Println(resultat)

}
