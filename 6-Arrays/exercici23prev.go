package main

import "fmt"

func main() {
	//Realitza un map de tres alumnes i tres notes.
	notes := map[string]int{ //nomVariable := map [tipus] valor{}
		"Cris":     9, //Clau Valor: String i valor
		"laura":    8, //Clau Valor: String i valor
		"Patricia": 8, //Clau Valor: String i valor. Aquesta última clau: valor ha de contenir també una coma!
	}

	fmt.Println(notes)
	fmt.Println(notes["Cris"]) //Per recuperar un valor en concret
	if ok:= (notes ["Cris"]) > 5:{
		fmt.Println ("Estàs aprovada!")
		} else {
		fmt.Println8("Estàs suspesa")
		}
	}
	

