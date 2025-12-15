package main

import "fmt"

func main() {

	//Programa para deternminar si eres mayor de edad o no
	var edad uint8
	fmt.Println("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")	
	}
}