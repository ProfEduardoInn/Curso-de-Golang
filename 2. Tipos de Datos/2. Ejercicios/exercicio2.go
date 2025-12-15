package main

import "fmt"

/*
Escribe para que sirven las distintas expresiones:
== != < > <= >=
Use las expresiones de arriba para comprobar que funcionan e imprima el resultado.
 - Pídele la edad al usuario y que al momemnto de imprimirlo uses los resultados de arriba e imprimas su resultado comparándolo com 18.
*/

func main() {

	fmt.Println("\"==\" Igual que: sive para comparar dos valores y ver si son iguales.")
	fmt.Println("\"!=\" Diferente: sive para comparar dos valores y ver si son distintos.")
	fmt.Println("\"<\" Menor que: sive para comparar dos valores y ver si uno es menor que el otro.")
	fmt.Println("\">\" Mayor que: sive para comparar dos valores y ver si uno es mayor que el otro.")
	fmt.Println("\">=\" Mayor o Igual que: sive para comparar dos valores y ver si uno es mayor o igual que el otro.")
	fmt.Println("\"<=\" Menor o Igual que: sive para comparar dos valores y ver si uno es mayor o igual que el otro.")

	var edad uint8
	fmt.Println("Cuál es tu edad? ")
	fmt.Scanf("%d", &edad)

	//fmt.Printf("%d == 18", edad)

	fmt.Printf("tu edad es %d\n", edad)
	fmt.Printf("tu edad %d es diferente de 18? %t\n", edad, edad != 18)
	fmt.Printf("tu edad %d es igual que 18? %t\n", edad, edad == 18)
	fmt.Printf("tu edad %d es mayor que 18? %t\n", edad, edad > 18)
	fmt.Printf("tu edad %d es menor que 18? %t\n", edad, edad < 18)
	fmt.Printf("tu edad %d es mayor o igual que 18? %t\n", edad, edad >= 18)
	fmt.Printf("tu edad %d es menor o igual que 18? %t\n", edad, edad <= 18)
}