package main

import "fmt"

var nombre string = "Pedro"
var edadNombre uint8 = 30

func main() {
	comidaFav := "Picanha \"pepp\" s "
	lugarFav := `Mi casa 
	es				
	mi lugar 			"favorito"
`
	fmt.Println(nombre)
	fmt.Printf("%T\n", nombre)
	fmt.Println(edadNombre)
	fmt.Printf("%T\n", edadNombre)
	fmt.Println(comidaFav)
	fmt.Printf("%T\n", comidaFav)
	fmt.Println(lugarFav)
	fmt.Printf("%T\n", lugarFav)
	//Como es un lenguaje de tipado estático, no se puede cambiar el tipo de dato
}