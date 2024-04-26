package main

import (
	"fmt"
	"strings"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

type item struct {
	produtoId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	item   []item
}

// gets e Setters (métodos)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p *pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p pessoa) setNomeCOmpleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	//ou
	produto2 := produto{"caderno", 2.5, 1}
	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2)
	//Aninhado

	pedido := pedido{
		userId: 1,
		item: []item{
			{1, 2, 12.10}, // slice
			{2, 1, 23.49},
			{11, 100, 3.14},
		},
	}

	fmt.Println(pedido)
	fmt.Println(pedido.valorTotal())

	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCOmpleto("Andre de jones")
	fmt.Println(p1)

}

//Médoto funcao com receiver (receptor)

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.item {
		total += item.preco * float64(item.qtde)
	}
	return total
}
