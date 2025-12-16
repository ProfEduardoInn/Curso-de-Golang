package main

import "fmt"

func main() {

	hola()
	holaTexto("Eduardo")
	adios()
	holaTexto("Andrés")
	adiosNumero(10)
	fmt.Println("Tu suma da:", sumaNumeros(2, 4))

}

func hola() {
	fmt.Println("Hola desde función hola")
}

func holaTexto (texto string) {
	fmt.Println("Hola", texto, "desde función holaTexto")
}

func adios() {
	fmt.Println("Hola desde función adiós")
}

func adiosNumero(numero int) {
	fmt.Println("Adios", numero)
}

func sumaNumeros(numero1, numero2 int) int {
	return numero1 + numero2
}

