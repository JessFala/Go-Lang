package main //Aquí és a on anomenem i es fa referència el paquet principal de Go

import ( //És opcional però acostuma a ser necessari el ús d’aquesta sintaxis a on
	"fmt" // invoquem els diferents mòduls que intervindran en el nostre script.
) //En el nostre cas hem d’emprar el fmt com a funció de sortida.

func main() { //La funció main acompanya el package i només es defineix un cop, aglutinant
	//la sintaxis principal del script.
	fmt.Println("Hola classe!")
}
