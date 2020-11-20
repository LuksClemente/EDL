package main

import "fmt"

func main(){
	/*Arrays
	1, 1, 2, 3, 5, 8, 13*/
	var n int
	var numeros [7]int
	for i := 0; i < 7; i = i + 1{
		fmt.Scan(&n)
		numeros[i] = n
	}
	fmt.Println(numeros)
}
