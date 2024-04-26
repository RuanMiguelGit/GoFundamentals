package main

import "fmt"

func main() {
	p := produto{"angelica", 1}
	p.toString()

	var pessoa imprimivel = produto{"Maria", 2}
	pessoa.toString()
}

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces sao implementadas implicidademente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome

}

func (p produto) toString() string {
	return fmt.Sprintf("%s %f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
