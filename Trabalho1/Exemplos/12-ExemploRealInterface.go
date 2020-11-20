// Golang program illustrates how 
// to implement an interface 
package main 

import "fmt"

// Creating an interface 
type tank interface { 

	// Methods 
	Tarea() float64 
	Volume() float64 
} 

type myvalue struct { 
	radius float64 
	height float64 
} 

// Implementing methods of 
// the tank interface 
func (m myvalue) Tarea() float64 { 

	return 2*m.radius*m.height + 
		2*3.14*m.radius*m.radius 
} 

func (m myvalue) Volume() float64 { 

	return 3.14 * m.radius * m.radius * m.height 
} 

// Main Method 
func main() { 

	// Accessing elements of 
	// the tank interface 
	var t tank 
	t = myvalue{10, 14} 
	fmt.Println("Area of tank :", t.Tarea()) 
	fmt.Println("Volume of tank:", t.Volume()) 
} 


//o codigo acima cria uma interface chamada tank, que é um conjunto dos métodos Tarea() e Volume(), ambos retornando um valor do tipo float64. O codigo escrito acima pode ser facilmente reutilizado apenas implementando a interface no decorrer do programa, e simulando uma orientação a objeto em go, já que permite que ocorra a criação de interfaces "filhas" ao criar interfaces que utilizam os metodos das interfaces "pais"
