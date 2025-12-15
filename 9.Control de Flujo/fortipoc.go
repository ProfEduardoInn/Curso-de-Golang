package main

import "fmt"

func main() {
	//Imprimir en pantalla los números del 1 al 10

	fmt.Println("1,2,3,4,5,6,7,8,9,10")

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//Imprimir los números del 10 al 1
	for x := 10; x >= 1; x-- {
		fmt.Println(x)
	}
	for i := 1; i <= 100; i+=2 {
		fmt.Println(i)
	}
	for i := 1; i <= 100; i+=5 {
		fmt.Println(i)
	}
	for i := 100; i >= 1; i-=2 {
		fmt.Println(i)
	}
	for i := 100; i >= 1; i-=5 {
		fmt.Println(i)		
	}
	for i := 1992; i <= 2025; i++ {
		fmt.Println(i)
	}
}