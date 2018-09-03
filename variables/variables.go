package main

import (
	"fmt"
)

/* Soma dois numeros*/
func sum(a, b int) int {
	return a + b
}

func main () {
	/* Declaração da variavel com tipo explicito */
	var result int
	result = sum(1, 2)
	
	/* Declaração com tipo implicito */
	var result2 = sum(2, 3)
	
	/* Declaração de forma curta */
	result3 := sum(3, 4) 
	
	fmt.Println("Resultado 1: ", result)
	fmt.Println("Resultado 2: ", result2)
	fmt.Println("Resultado 3: ", result3)
}
