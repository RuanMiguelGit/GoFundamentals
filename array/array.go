package main

import "fmt"

func main() {
	fmt.Println(arr())
}

func arr() float64 {
	//homogernia, estatico e com valor fixo

	var nota [3]float64
	nota[0], nota[1], nota[2] = 1, 2, 3
	fmt.Println(nota)

	total := 0.0

	for i := 0; i < len(nota); i++ {
		total += nota[i]
	}
	return total
}
