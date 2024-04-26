package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d\n)", pessoa, texto, i+1)
	}
}

func main() {
	go fale("Andre", "Pq vc nao fala comigo", 3)
	fale("Joao", "Só pode falar depois de vc", 1)

}
