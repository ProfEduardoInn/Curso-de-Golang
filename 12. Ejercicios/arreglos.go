package main

import "fmt"

func main() {
	//Crea un arreglo que contenga 5 números enteros
	//Pídelos al usuario
	//Imprime los valores con su índice

	arreglo := [5]int{}

	for i := 0; i < len(arreglo); i++ {
		fmt.Print("Dame el valor #", (i+1), ": ")
		fmt.Scan(&arreglo[i])
	}
	fmt.Println(arreglo)

	for i, v := range arreglo {
		fmt.Println("indice:", i, "Valor:", v)
	}
	
	for i := 0; i < len(arreglo); i++ {
		fmt.Println("indice:", i, "Valor:", arreglo[i])
	}
}