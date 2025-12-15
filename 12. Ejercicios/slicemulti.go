package main

import "fmt"

func main() {

	/*
	- Crea un slice de dos dimensiones del tipo uint16.
	- Agrégale los años de tus amigos.
	- Los datos los debe ingresar el usuario.
	- Que se vea de manera gráfica así:
		[[123, 456, 789 ]]
		[[234, 567, 890]]
	- Imprime el índice y el valor de los elementos del slice.
	*/

	/*var y1y2[][] uint16
	fmt.Println(y1y2)*/

	year1 := []uint16{0, 0, 0}
	year2 := []uint16{0, 0, 0}

	y1y2 := [][]uint16{year1, year2}

	for i := 0; i < 2; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Println(j)
			fmt.Println(y1y2[i][j])
			fmt.Println("Ingresa el año: ")
			fmt.Scan(&y1y2[i][j])
			fmt.Println("Agregaste el año", y1y2[i][j])
		}
	}
	fmt.Println(y1y2)
	
	for i := 0; i < len(y1y2); i++ {
		fmt.Print("[[")
        for j := 0; j < len(y1y2[i]); j++ {
				fmt.Print(" ", y1y2[i][j], " ")
			
        }
		fmt.Print("]]")
		fmt.Println()
	}

	for indice, valor := range y1y2 {
		fmt.Println("Índice: ", indice, " y su valor es: ", valor)
		for i, v := range valor {
			fmt.Println("Índice: ", i, " y su valor es: ", v)
		}
	}

	for _, valor := range y1y2 {
    fmt.Println("[", valor, "]")
	}
}