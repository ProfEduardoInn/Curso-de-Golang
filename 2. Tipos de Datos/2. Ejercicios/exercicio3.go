package main

import "fmt"

/*

 - Crea constantes con un tipo de dato específico y sin tipo de dato específico. 
 - Explica cuales son las diferencias entre ambos y cuando usar que tipo de constantes.
 - Imprime el valor de las constantes con su tipo de dato.
 - También su dirección de memoria.

*/

func main() {

	const (
		x = 22 
		y string = "Hola"
	)
	xp := x
	yp := y

	fmt.Println("Las constantes sin tipo son flexibles, pueden ser de cualquier tipo, aptas para cálculos numerícos, las constantes con tipo son rígidas, se usan cuando se requiere un valor exacto, para proyectos grandes.")

	fmt.Println()
	fmt.Printf("- El valor de la constante \"x\" es: %v\n- Es una constante sin tipo específico\n- Su tipo de dato es: %T\n- Su dirección de memoria es: %p", x, x, &xp)

	fmt.Println()
	fmt.Printf("\n- El valor de la constante \"y\" es: %v\n- Es una constante con tipo específico\n- Su tipo de dato es: %T\n- Su dirección de memoria es: %p", y, y, &yp)

	fmt.Println()
	fmt.Println("\nRealmente no se puede imprimir la dirección de memoria de una constante, se necesita declarar una variable y asignarle el valor de la constante para después si imprimir la dirección de la variable.")
}