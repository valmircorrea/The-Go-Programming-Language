package main 

import (
	"fmt"
)

func maior(a, b int) int {
	if (a > b) {
		return a
	} else {
		return b;	
	}
} // Não precisa de 'return' nas funções

func main() {
	a := 5
	b := 4
	fmt.Printf("Maior entre %d e %d: %d", a, b, maior(a, b))
}
