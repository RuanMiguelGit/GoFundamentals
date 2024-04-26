package main

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F48"
	f.velocidadeAtual = 0
	f.turboLigado = true
}

//
