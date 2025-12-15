package main

import "fmt"

func main() {

	edad := 0

	fmt.Printf("Presiona 0 para salir del programa\n")

	for {
		
		fmt.Printf("Dame tu edad: ")
		fmt.Scanf("%d", &edad)
	
		if edad >= 18 {
			fmt.Printf("Tienes %d año(s). Eres mayor de edad.\n", edad)
		} else if edad == 17 {
			fmt.Printf("Tienes %d año(s). Si me das dinero, te dejo pasar.\n", edad)
		} else {
			fmt.Printf("Tienes %d año(s). Eres menor de edad.\n", edad)
		}
		if edad == 0 {
			break
		}
	}
}