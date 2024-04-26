package main

import "fmt"

func main() {
	fmt.Println(retornoNomeado(1, 2))
	x, y := retornoNomeado(3, 6)
	fmt.Println(x, y)
	fmt.Println(media(1, 4, 67, 8, 9, 0))

	arr := []float64{1, 2, 3}
	fmt.Println(media(arr...))
	obterValorAprovado()
}

func retornoNomeado(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

// funcoes variaveis

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

//**********************************************************************

//defer

func obterValorAprovado() {
	defer fmt.Println("fim")
	fmt.Println("inicio")
}

//PAssando ponteiros

func pont(n *int) {
	*n++
}

//funcao init
func init() {
	fmt.Println("INICIO")
}
