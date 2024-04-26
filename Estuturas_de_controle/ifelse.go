package main

import "fmt"

func main() {
	notasCondic(7.8)
	fmt.Println((notasCondicIf(7)))

}

func notasCondic(nota float64) {
	if nota >= 7 {
		println("Aprovado")
	} else {
		println("Reprovado")
	}
}

func notasCondicIf(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 4 && n < 9 {
		return "B"
	} else if n >= 0 && n < 4 {
		return "C"
	} else {
		return "Nota inválida"
	}
}
