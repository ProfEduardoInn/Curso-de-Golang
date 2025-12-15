package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	/*
	- agrega el elemento 11 al final del slice.
	- imprime el slice.
	- agrega el valor del slice a un nuevo slice con los valores agregados de 12, 13, 14, 15.
	- Imprime el nuevo slice.
	*/

	slice = append(slice, 11)
	fmt.Println(slice)
	
	slice2 := append(slice, 12, 13, 14, 15)
	fmt.Println(slice2)

	sliceTwo := []int{12, 13, 14, 15}
	slice = append(slice, sliceTwo...)
	fmt.Println(slice)

}