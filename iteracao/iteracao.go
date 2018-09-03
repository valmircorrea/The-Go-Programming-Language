package main 

import (
	"fmt"
)

func main() {
	/* nao fecha com parenteses*/
	for ii := 0; ii < 10; ii++ {
		fmt.Printf("%d ", ii)
	}
	
	fmt.Println(" ")
	
	/* Fazer um while */
	pp := 1
	for pp < 10 {
		fmt.Printf("%d ", pp)
		pp += pp
	}
}
