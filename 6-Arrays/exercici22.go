package main

import (
	"fmt"
)

func main() {
	x := make([]int, 9, 10)

	x[2] = 3
	x[6] = 7

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x[11] = 11 //Això donarà error perquè només existeix la capacitat fins a 10.
	x = append(x, 11) // A l'afegir l'11 al slice doblarà la capacitat i ara serà de 20.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //Això duplicarà l'espai de memòria

}
