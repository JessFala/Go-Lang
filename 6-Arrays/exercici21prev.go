package Main

import "fmt"

func main() {
	treballadorxs := []string{"Josep", "Cristina", "Helena", "Robert"} //1r Slice

	novesIncorporacions := []string{"Ivan", "Estel", "Aitana"} //2n Slice

	//Ara has de juntar les dues estructures en una de sola.

	treballadorxs = append(treballadorxs, novesIncorporacions...)
	fmt.Println(treballadorxs)
}
