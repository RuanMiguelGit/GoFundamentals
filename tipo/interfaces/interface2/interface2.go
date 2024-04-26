package main

import "fmt"

func main() {
	carro1 := ferrari{"F40", 1, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1)

	var carro2 esportivo = &ferrari{"F40", 2, 0}
	fmt.Println(carro2)

}

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo             string
	velocidadeAnterior int
	velocidadeAtual    int
}

func (f *ferrari) ligarTurbo() {
	f.ligarTurbo()
}
