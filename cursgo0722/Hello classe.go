// You can edit this code!
// Click here and start typing.
package main


import (
	"fmt"
 ) //Llibreria (abreviatura de "format")
/* El "." és l'element de concatenació. La primera lletra del mètode ('Print') va en Majúscules;
vol dir que aquest mètode és accessible = visible (vs. no accessible = no visible) */

func main() { /* Aquí NO s'utilitzen els delimitadors de sentència de programació així que no cal acabar amb ';'. */
	fmt.Println("Hello, classe!!!") //Sempre cometes dobles!!!
}

/* Assignació (fixa) el valor d'una variable mitjançant 
un operador abreujat (:=) i a més dir el tipus de dada que és. 
El nom de la variable NO pot contenir un número a l'inici (sí al final), 
ni caràcters especials (accents, ñ, ...)
GO NO admet una variable orfe. No farà córrer el codi, no pot compilar. */




 /*## Assignació múltiples ## 
 NO cal que siguin dades del mateix tipus. */

j, k, l := "shark", 2.05, 15
fmt.Println(j)

## Concatenació ## 
func main() {
	nom := "Jéssica"
	fmt.Println(nom)
	fmt.Println("Hola classe! El meu nom és ", nom)  //Concatenar ("text/variable1", variable2)"
}

## Agrupació de variables ##
var ( 
	casa = "unifamilar" 
	color = "verd" 
)

/*Perquè permeti compilar malgrat no li donem una de les variables (ex. salutació) */

salutació, err:=  // declaració tradicional correcta
_, err:=  // El (_) salva el no ús de la variable 'salutació'


## Constants ##
/* És només de lectura, però NO es modifiquen*/
/* Agrupar constants: NO comes (,) i NO punt i coma final (;)*/

const ( 
	tenda1 = “CiberShop” 
	ciutat = “Barcelona”
 )


