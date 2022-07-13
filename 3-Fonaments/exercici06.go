package main

import (
	"fmt"
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a) //4    100
	// \t imprimirá una tabulació i %b mostrara en binari
	//Hem de imprimir un binari amb una posició més cap a la dreta
	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b) //8    1000
}
