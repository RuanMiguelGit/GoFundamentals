package main

import "fmt"

func main() {
	//map = objetos
	//mapas devem ser inicializados
	// var aprovados map[int]string
	aprovados := make(map[int]string)
	aprovados[1] = "maria"
	aprovados[2] = "joao"
	aprovados[3] = "Ana"

	// fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)
	}

	// fmt.Println(aprovados[2])
	delete(aprovados, 2)
	// fmt.Println((aprovados))

	//*******************************************
	funcAp := map[string]float64{
		"joao":  1,
		"maria": 2,
	}

	fmt.Println(funcAp)
	//...........................................................................
	//Maps aninhados

	mapAninhado := map[string]map[string]float64{
		"G": {
			"Gabriela": 123,
		},
		"P": {
			"Andreaia": 234,
		},
	}
	fmt.Println(mapAninhado)
}
