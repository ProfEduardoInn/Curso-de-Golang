package main

import "fmt"

func main() {
	/*Imprime todos los números del 1 al 10 cuyo resuduo es igual a cero si es dividido entre 2.

	Imprime todos los números del 1 al 100 múltiplos de 3, pero no los múltiplos de 5.

	Imprime todos los números del 1 al 100 múltiplos de 2, pero no los múltiplos de 3.

	Imprime todos los números del 1 al 100 divisibles entre 3 y 5.

	Imprime todos los números del 1 al 100 divisibles entre 2 y 3.
    */


	fmt.Println("Números divididos por 2 con resto 0 entre 1 y 10 ")
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Printf("[%d] ", i)
		}
	}

	fmt.Println("\nNúmeros múltiplos de 3, pero no de 5")
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Printf("[%d] ", i)
		}
	}
	
	fmt.Println("\nNúmeros múltiplos de 2, pero no de 3")
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 && i % 3 != 0 {
			fmt.Printf("[%d] ", i)
		}
	}
	
	fmt.Println("\nNúmeros divisibles por 3 y 5")
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("[%d] ", i)
		}
	}
	
	fmt.Println("\nNúmeros divisibles por 2 y 3")
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 && i % 3 == 0 {
			fmt.Printf("[%d] ", i)
		}
	}
	
	fmt.Println("\nResiduos de 3 y 5")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		fmt.Printf("El residuo de 3: %d, residuo de 5: %d\n", i % 3, i % 5)
	}
	
	fmt.Println("\nResiduos de 2 y 3")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		fmt.Printf("El residuo de 2: %d, residuo de 3: %d\n", i % 2, i % 3)
	}
}