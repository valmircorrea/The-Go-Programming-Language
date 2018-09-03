package main 

import (
    "fmt"
)

/**
 * Soma de dois numeros (na definição, o tipo vem depois)
 */
func add(x int, y int) int {
	return x + y
}

/**
 * Soma de três numeros (Quando os parametros sao do mesmo tipo, pode omitir todos e definir no ultimo)
 */
func sum(x , y, z int) int {
	return x + y + z
}

func main () {
	
	var resultAdd = add(2, 3)
	var resultSum int = sum(2, 2, 2)
	
	fmt.Println("Soma de dois numeros: ", resultAdd)
	fmt.Println("Soma de três numeros: ", resultSum)
}
