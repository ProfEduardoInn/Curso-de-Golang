package main

import "fmt"

func main() {

	var edad uint8
	fmt.Println("Ingresa tu edad")
	fmt.Scanf("%d", &edad)

/*	if edad == 10 {
		fmt.Println("La edad es igual a 10")
	} else if edad == 11 {
		fmt.Println("La edad es igual a 11")
	} else if edad == 12 {
		fmt.Println("La edad es igual a 12")
	} else if edad == 13 {
		fmt.Println("La edad es igual a 13")
	} else if edad == 14 {
		fmt.Println("La edad es igual a 14")
	} else if edad == 15 {
		fmt.Println("La edad es igual a 15")
	} else {
		fmt.Println("La edad es diferente de 10 o 15")
	}
	*/

	//Switch
	switch edad {
	case 10:
		fmt.Println("La edad es 10")
	case 11:
		fmt.Println("La edad es 11")
	case 12:
		fmt.Println("La edad es 12")
	case 13:
		fmt.Println("La edad es 13")
	case 14:
		fmt.Println("La edad es 14")
	case 15:
		fmt.Println("La edad es 15")
	default:
		fmt.Println("La edad no está entre 10 o 15")
	}
}