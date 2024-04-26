package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(swithCase(10))
	switchTwo()

}

func swithCase(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	default:
		return "Nota Invalida"
	}
}

func switchTwo() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde;")
	case t.Hour() < 24:
		fmt.Println("Boa Noite")
	default:
		fmt.Println("Hora invalida")
	}
}
