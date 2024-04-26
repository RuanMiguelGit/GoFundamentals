package main

import "fmt"

func main() {
	var e esportivoLuxooso = bmw7{}

	e.ligarTurbo()
	e.fazerBaliza()

	//******************************

	var coisa interface{}
	coisa = 2
	coisa = curso{"curso teste"}

	fmt.Println(coisa)
	//consguimos criar uma interface vazia e ir implicitamente atribuindo valores a ela

}

type curso struct {
	curso string
}

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxooso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("esportivo")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("baiza")
}

//*****************************************************************************
