package main

import "fmt"

func calcular(a int ) ( int , int ){
	var quadrado int = a*a
	var cubo int = a*a*a
	
	return quadrado, cubo
}
func main(){
	var x int
	
	fmt.Print("Digite um valor para x: ")
	fmt.Scan(&x)
	fmt.Println("Quadrado e cubo:", calcular(x))
}
