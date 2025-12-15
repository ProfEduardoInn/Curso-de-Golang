package main

import "fmt"

func main() {
	
	//Crea un arreglo que contenga 5 números enteros
	//Pídelos al usuario
	//Imprime los valores con su índice

	//var slice []int
	
	slice := make([]int, 5, 10)

	for i := 0; i < len(slice); i++ {
		fmt.Print("Dame el valor #", (i+1), ": ")
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice)

	for i, v := range slice {
		fmt.Println("indice:", i, "Valor:", v)
	}
	
	for i := 0; i < len(slice); i++ {
		fmt.Println("indice:", i, "Valor:", slice[i])
	}

	slice2 := []int{0, 0, 0, 0, 0}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	for i := 0; i < len(slice2); i++ {
		fmt.Print("Dame el valor #", (i+1), ": ")
		fmt.Scan(&slice2[i])
	}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	for i, v := range slice2 {
		fmt.Println("indice:", i, "Valor:", v)
	}
	
	for i := 0; i < len(slice2); i++ {
		fmt.Println("indice:", i, "Valor:", slice2[i])
	}

	fmt.Println(slice2)
}