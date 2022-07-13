package Main

import (
	"fmt"
)

//Afegeix al carro de la compra "Oli d'oliva"

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}

	carro[1] = "cogombre"
	carro[2] = "Fanta" /*Ens marca error si no li afegim un valor perquè Go entén que no l'estem utilitzant,
	per això hem substituït 'aigua' per 'Fanta'.*/

	carro = append(carro, "Oli d'oliva") //Carro = Slice + ""(nou element afegit a l'array; en aquest cas "oli d'oliva")
	fmt.Println(carro)                   /* La 'P' de print en majúscula és per indicar que és visible)*/
}
