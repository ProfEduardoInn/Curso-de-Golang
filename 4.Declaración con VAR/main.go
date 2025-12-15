package main

import "fmt"

//esta si se puede usar fuera de las funciones
//variable global, lo ideal es usarlo dentro de la función y de la forma corta
var edad uint8
var feliz bool
var flotante float32


func main() {
	//declara una variable y asigna el valor al mismo tiempo

	// ! Operador de declaración corta
	// No se puede usar fuera de una función
	nombre := "Eduardo"
	fmt.Println(nombre, edad)
	edad = 20
	fmt.Println(edad)
	fmt.Println(feliz)

	var apellido string = "Rojas"
	var comidaFav string
	feliz = true
	fmt.Println(nombre, apellido, edad)
	comidaFav = "Picanha"
	fmt.Println(nombre, apellido, edad, comidaFav, feliz, flotante)
	printaEdad()

}

func printaEdad() {
	fmt.Println("Mi edad es: ", edad)
}