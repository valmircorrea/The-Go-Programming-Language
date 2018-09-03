package main

/*
 ==============================================================================
 ARQUIVO............: corridaDeRevezamento.go
 DESCRICAO..........: Programa que simula uma corrida de revezamento em que uma
					   equipe possui quatro corredores. Utilizado um canal de
					   comunicação sem buffer, para promover sincrinização.
 AUTOR..............: Valmir Correa da Silva Junior (valmir.correa@ufrn.edu.br)
 CRIADO EM..........: 06/2018
 MODIFICADO EM......: 06/2018
 ===============================================================================
*/

import (
	"fmt"
	"time"
)

// NumCorredores - Numero máximo de corredores
const NumCorredores int = 4

func main() {
	/* Criando um canal sem buffer */
	bastao := make(chan int)

	/* Cria o primeiro corredor com a Go Rotina */
	go Corredor(bastao)

	/* Começa a corrida */
	fmt.Printf("Começa a corrida!\n")
	bastao <- 1

	/* Tempo para a corrida */
	time.Sleep(15 * time.Second)
}

// Corredor - O Corredor
func Corredor(bastao chan int) {

	/* Recebimento do bastao pelo canal */
	corredor := <-bastao
	fmt.Printf("Corredor %d está correndo com o bastão!\n", corredor)

	/* Proximo corredor se prepara para continuar a corrida */
	var proxCorredor int
	if corredor != NumCorredores {
		proxCorredor = corredor + 1
		preparaParaReceberOBastao(proxCorredor, bastao)
	}

	/* Tempo para cada corredor fazer o percurso */
	time.Sleep(2 * time.Second)

	/* Fim da corrida */
	if corredor == 4 {
		fmt.Printf("Corredor %d completou o percurso! Corrida finalizada!\n", corredor)
		return
	}

	/* Passa o bastao para o proximo corredor */
	fmt.Printf("Corredor %d passa o bastao para o Corredor %d!\n", corredor, proxCorredor)
	bastao <- proxCorredor
}

/* Prepara o proximo corredor */
func preparaParaReceberOBastao(proxCorredor int, bastao chan int) {
	/* Tempo para a passagem do bastao */
	time.Sleep(2 * time.Second)

	fmt.Printf("Corredor %d preparado para recerber o bastão do corredor %d!\n", proxCorredor, proxCorredor-1)
	go Corredor(bastao)
}
