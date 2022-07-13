package Main

import (
	"fmt"
)

func main() {
	notes := map[string]int{
		"Pablo":     5,
		"Enriqueta": 9,
		"Jordi":     4,
	}

	fmt.Println(notes)
	fmt.Println(notes["Enriqueta"])

}
