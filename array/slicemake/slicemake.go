package main

import "fmt"

func main() {
	arr := make([]int, 10)
	arr[9] = 12
	fmt.Println(arr)

	s := make([]int, 10, 20)
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, cap(s))
}
