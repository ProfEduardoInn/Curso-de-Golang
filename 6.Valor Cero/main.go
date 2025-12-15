package main

import "fmt"

var nombre string
var edad int
var peso float32
var tarea bool
var complejo complex128

func main() {
	fmt.Println(nombre)
	fmt.Printf("%T, %v, %#v, %##\n", nombre, nombre, nombre, nombre)
	fmt.Println(edad)
	fmt.Printf("%T\n", edad)
	fmt.Println(peso)
	fmt.Printf("%T\n", peso)
	fmt.Println(tarea)
	fmt.Printf("%T\n", tarea)
	fmt.Println(complejo)
	fmt.Printf("%T\n", complejo)
}