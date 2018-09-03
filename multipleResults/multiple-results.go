package main

import "fmt"

/* Função pode ter multiplos retornos*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	
	/* Usando ':=', é a forma curta de declarar variavel, sem o 'var' */
	a, b := swap("Hello", "World")
	
	fmt.Println(a, b)
}
