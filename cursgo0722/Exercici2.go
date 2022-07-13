package main

import "fmt"

func Main() {
	var a int
	var b string
	var c float64
	var d bool //Per defecte s'entén com a "False".

	fmt.Printf("var a %T = %+v\n", a, a) // (\n) Fa salt de línia
	fmt.Printf("var a %T = %+v\n", b, b)
	fmt.Printf("var a %T = %+v\n", c, c)
	fmt.Printf("var a %T = %+v\n", d, d)
}

func main() {
	nom := "Jéssica"
	fmt.Println(nom)
	fmt.Println("Hola classe! El meu nom és ", nom) //Concatenar
}
