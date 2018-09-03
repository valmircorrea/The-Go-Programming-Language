package main 

import (
	"fmt"
)

/* Criando o "Student" como ALIAS para um tipo string */
type Student string 

/* Definindo os valores para um carro */
func defStudent() Student {
	
	var student Student
	
	student = "Valmir"
	
	return student
}

/* Criando uma estrutura que define um carro */ 
type Car struct {
	name string
	ano int
}

/* Definindo os valores para um carro*/
func defCar() Car{
	var car Car
	
	car = Car{"Corolla", 2018}
	
	return car
}

func main() {
	
	fmt.Println( defStudent() )
	fmt.Println( defCar() )
}

