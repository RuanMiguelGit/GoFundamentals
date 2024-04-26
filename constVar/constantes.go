package main

import ("fmt"; "math")

func main() {
	const PI float64 = 3.45
	var raio = 3.2

	//forma reduzida de criar uma var

	area := PI * math.Pow(raio, 2)

	const ( a =1, b=2)
	var e,f bool =true, false
	g, h , i := 2, false, "epa"
	fmt.Println("area é", area)
}