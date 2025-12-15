package main

import "fmt"

func main() {
	
	/*
	- Crea un programa que diga que eres mayor de edad.
	
	- Si eres mayor de edad que diga: eres mayor de edad, si eres menor de edad que diga: eres menor de edad.
	*/

	edad := 0

	fmt.Printf("Presiona 0 para salir del programa\n")
	for {
		
		fmt.Printf("Dame tu edad: ")
		fmt.Scanf("%d", &edad)
	
		if edad >= 18 {
			fmt.Printf("Tienes %d año(s). Eres mayor de edad.\n", edad)
		} else {
			fmt.Printf("Tienes %d año(s). Eres menor de edad.\n", edad)
		}
		if edad == 0 {
			break
		}
	}
}