package main

import "fmt"

func main() {
	//% Símbolo módulo
	fmt.Println("Números divisibles en 2 entre 0 y 100")
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Números divisibles en 3 entre 0 y 100")
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}