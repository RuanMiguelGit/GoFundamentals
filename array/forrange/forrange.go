package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	for i, numero := range arr {
		fmt.Println("numero", i, numero)
	}
}
