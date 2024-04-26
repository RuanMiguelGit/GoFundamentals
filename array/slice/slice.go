package main

import "fmt"

func main() {
	// a1 := [3]int{1, 2, 3} //array
	// s1 := []int{1, 2, 3}  // slice, tamanho variavel

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice nao é um array, slice deine um pedaco do array

	s2 := a2[1:3]
	s3 := a2[:2]
	fmt.Println(s2, a2)
	fmt.Println(s3)

}
