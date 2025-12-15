package main

import "fmt"

func main() {
	//Operador !
	//Convierte True en False y viceversa
	a := 10
	b := 5
	if !(a < b) {
		fmt.Println("a es menor que b")
	}

	if true {
		fmt.Println("Verdadero normal")
	}
	if false {
		fmt.Println("Falso normal")
	}
	if !true {
		fmt.Println("Verdadero negado")
	}
	if !false {
		fmt.Println("Falso negado")
	}
	if 2 == 2 {
		fmt.Println("2 es igual a 2")
	}
	if 2 != 2 {
		fmt.Println("2 es menor o igual a 2")
	}
	if edad := 20; edad > 18 {
		fmt.Println("Mayor de edad")
	}
}