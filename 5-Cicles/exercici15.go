//Dibuixa un cuadrat amb un únic * i que tingui 10 unitats d'ample i 10 d'alt.

package main

import (
	"fmt"
)

func main() {
	c := "*"
	for i := 0; i < 10; i++ {
		fmt.Println(c)
		for j := 0; j < 10; j++ {
			fmt.Println(c)
		}
	}
}
