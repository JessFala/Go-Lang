/*Hem de realitzar una estructura condicional composta que contempli una nota y imprimeixi el següent:
- Menor de 5 "Suspes"
- Cinc "aprobat"
- Entre 6 i 8 (inclós) "Notable"
- En qualsevol altre cas "Excelent"
*/

package main

import (
	"fmt"
)

func main() {
	nota := 8

	if nota < 5 {
		fmt.Println("Suspes")
	} else if nota < 6 {
		fmt.Println("Aprobat")
	} else if nota < 9 {
		fmt.Println("Notable")
	} else {
		fmt.Println("Excelent")
	}
}
