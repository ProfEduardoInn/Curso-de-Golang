package main

import "fmt"

func main() {
	/*
	- Crea un mapa con tus contactos que tenga nombre y teléfono.
	- Para cada contacto imprime el nombre y teléfono.
	! Guarda más de dos números por persona
	- Imprime el índice y el valor del mapa
	- Pista: usa un slice para guardar los teléfonos.
	*/

	contactos := map[string][]uint32 {
		"Bernardo": {31997848, 1234567, 31254},
		"Jessé": {3199799, 24567895},
		"Arnaldo": {3899988, 54897894, 6644, 521},
		"Osvaldo": {35998745, 34567891},
	}


	fmt.Println(contactos)

	for i, v := range contactos {
		fmt.Println("Nombre:", i)
		for i, valor := range v {
			fmt.Printf("Teléfono #%d: %d\n", (i+1), valor)
		}
	}

	for indice, valor := range contactos {
		fmt.Println("El índice es:", indice)
		for i, v := range valor {
			fmt.Printf("El valor #%d: %d\n", i, v)
		}
	}
}