package main

import "fmt"

func main() {
 fmt.Print("Linha")
 fmt.Println("Linha nova") 
 fmt.Println("Linha nova")

 x := 3

 xs := fmt.Sprint(x)

 fmt.Println("O valor de x é" + xs)
 fmt.Printf("O valor é %x", x)

}