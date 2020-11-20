package main

import "fmt"

func soma(x int , y int ) int{
	return x + y
}
func main(){
	var x, y, resultado int
	
	fmt.Print("Digite um valor para x: ")
	fmt.Scan(&x)
	fmt.Print("Digite um valor para y: ")
	fmt.Scan(&y)
	resultado = soma(x, y)
	fmt.Println("O resultado da soma é:", resultado)
}
	
