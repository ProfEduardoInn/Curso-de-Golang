package main

import "fmt"

func main() {
	const nombre = "Eduardo"
	fmt.Println("Mi nombre es: ", nombre)
	fmt.Println(nombre)
	fmt.Println()
	n, err := fmt.Println(nombre, true, 42, 42.10, "hola")
	// _ se usa para ignorar el error, se reemplaza err
	fmt.Println(n)
	fmt.Println(err)
}