package main

import (
	"fmt";
	"strings"
)

func main() {
	/*
	- Usando el mapa anterior, agrega un nuevo contacto con el nombre "Arnaldo" y los teléfonos "123456" y "567890". Borra el contacto anterior de "Arnaldo" antes de agregar el nuevo contacto. Imprime el mapa actualizado.

	*/
	
	contactos := map[string][]uint32 {
		"Bernardo": {31997848, 1234567, 31254},
		"Jessé": {3199799, 24567895},
		"Arnaldo": {3899988, 54897894, 6644, 521},
		"Osvaldo": {35998745, 34567891},
	}

	fmt.Println(contactos)

	var addEraseContact uint8
	for {
		fmt.Println("Deseas agregar o borrar un contacto? ")
		fmt.Println("Presiona 1 para agregar")
		fmt.Println("Presiona 2 para borrar")
		fmt.Println("Presiona 0 para salir")
		fmt.Scan(&addEraseContact)

		var newContact string
		var telef uint32
		var referen int
		var continuar string

		switch addEraseContact {
		case 1:
			for {
				var numberTelef []uint32
				fmt.Println("Dame un nuevo contacto: ")
				fmt.Scan(&newContact)

				fmt.Println("Cuántos telefonos quieres introducir en ese contacto?")
				fmt.Scan(&referen)

				for i := 0; i < referen; i++ {
					fmt.Println("Dame el número de telefono #", i+1)
					fmt.Scan(&telef)

					numberTelef = append(numberTelef, telef)	
				}

				contactos[newContact] = numberTelef

				fmt.Println("Deseas continuar agregando contactos?" )
				fmt.Scan(&continuar)
				if strings.ToLower(continuar) == "si" {

	 			} else { 
					break
				}
			} 
		case 2:
			for {

				var eraseContact string
				fmt.Println("Cual contacto quieres borrar? ")
				fmt.Scan(&eraseContact)

				if _, ok := contactos[eraseContact]; ok {
					delete(contactos, eraseContact)
        			fmt.Println("Contacto", eraseContact, " eliminado con éxito!")
				} else {
					fmt.Println("ERROR! el contacto ", eraseContact, "no existe!")
				}
				
				var contactDelete string
				fmt.Println("Deseas borrar más contactos? ")
				fmt.Scan(&contactDelete)
				if strings.ToLower(contactDelete) == "si" {
		
				} else {
					break
				}
			}
		case 0:
			fmt.Println("Has terminado de usar el programa.")
			fmt.Println("Estado Final de la lista de contactos")
			for i, v := range contactos {
				fmt.Println("Nombre:", i)
				for i, valor := range v {
					fmt.Printf("Tel %d: %d\n", i+1, valor)
				}
			}
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}