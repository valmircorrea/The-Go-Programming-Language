package main 

import (
	"fmt"
)

func main() {
	/* Definição curta e sem tipagem explicita*/
	num := 10
	
	/* switch case*/
	switch {
		case num < 0:
			fmt.Println("O numero é negativo")
		case num >= 0 && num <= 10:	
			fmt.Println("O numero esta entre 0 e 10")
		default: // default é opcional
			fmt.Println("O numero é maior que 10")
	}
}

