package main
import "fmt"

func soma(a int, b int) int  {
	return a + b 
}

func print(valor int) {
 fmt.Println(valor)
}

func main() {
	resultado := soma(3, 4)
	print(resultado)
}