package main

import "fmt"

func main() {

	/*for{
		fmt.Println("Este es un bucle infinito")
		if condicion {
			break
		}
	}*/

	for i := 0; i <= 10 ; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Printf("Valor de i: %v\nValor de J: %v\n", i, j)
		}
	}
	for i := 0; i <= 5; i++ {
		fmt.Printf("Bucle externo: %d \n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\tBucle interno: %d \n", j)
		}

		
	}
}