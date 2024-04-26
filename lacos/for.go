package main

import "fmt"

func main() {
	lacoFor()
	lacoForTrad()
}

func lacoFor() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

}

func lacoForTrad() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

}
