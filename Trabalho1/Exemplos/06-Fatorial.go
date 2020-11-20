  
package main

import "fmt"

func fat(x int ) int{
	if(x == 0){
		return 1
	}else{
		return x*fat(x - 1)
	}
}
func main(){
	var x int
	
	fmt.Print("Digite um valor para x: ")
	fmt.Scan(&x)
	fmt.Println("Fatorial:", fat(x))
}
