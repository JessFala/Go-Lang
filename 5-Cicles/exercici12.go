//Mostra els numeros parells del o al 20.
package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 20 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
