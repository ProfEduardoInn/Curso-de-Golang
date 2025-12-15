package main

import "fmt"

func main() {
	var dia string
	fmt.Println("Escribe salir para salir")
	
	for dia != "salir" {

		fmt.Println("dame el dia de la semana: ")
		fmt.Scan(&dia)

		switch dia {
		case "lunes":
			fmt.Println("es el día 2, Monday")
		case "martes":
			fmt.Println("es el día 3, Tuesday")
		case "miércoles":
			fmt.Println("es el día 4, Wednesday")
		case "jueves":
			fmt.Println("es el día 5, Thursday")
		case "viernes":
			fmt.Println("es el día 6, Friday")
		case "sábado":
			fmt.Println("es el día 7, Saturday")
		case "domingo":
			fmt.Println("es el día 1, Sunday")
		default:
			fmt.Println("No es un día de la semana")
			break
		}
	}

}
