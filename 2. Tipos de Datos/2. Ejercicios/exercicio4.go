package main

import "fmt"

/*

 - Escribe un programa que asigne valores a una variable de tipo entero, que el valor lo ingrese el usuario.
 - Imprime sus valores en pantalla (decimal, binario, octal y hexadecunal)
 - Usa el bitShifting de 1 a la izquierda y asigna ese valor a una nueva variable.
 - Imprime los valores en pantalla de la nueva variable (decimal, binario, octal y hexadecimal)
 - Explica que es el bitShifting y para qee sirve.

*/

func main() {

	var num int64
	fmt.Println("Ingrese un valor: ")
	fmt.Scanf("%v", &num)
	fmt.Printf("\nTu valor es: %v\n\n- En Decimal: %d\n- En Binario: %b\n- En Octal: %o\n- En hexadecimal: %x\n", num, num, num, num, num)

	fmt.Println("\nUsando el bitShifting")

	num2 := num << 1

	fmt.Printf("\nTu valor es: %v\n\n- En Decimal: %d\n- En Binario: %b\n- En Octal: %o\n- En hexadecimal: %x\n", num2, num2, num2, num2, num2)

	fmt.Println("\nEl bitShifting es usado para mover bytes de una variable en binario (base 2) de derecha a izquiereda o viceversa. Este cambia el valor de una variable.")

}
