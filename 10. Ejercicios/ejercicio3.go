package main

import "fmt"

func main() {

	/*
	Imprime los años que has estado vivo hasta el año actual
	Usando o bucle for condition {}
	Ussando o bucle for {}
	*/

	fmt.Println("For")
	for i := 1992; i <= 2025; i++ {
		fmt.Printf("%d\n", i)
	}
	
	fmt.Println("For tipo while")

	currentYear := 2025
	initYear := 1992
	countYear := 0
	
	for initYear <= currentYear {
		fmt.Printf("%d\n", initYear)
		initYear++
		countYear++
	}
	fmt.Printf("Has vivido %d años", countYear)
	
	fmt.Println("For tipo do-while")
	i := 1992
	for {
		if i > 2025 {
			break
		}
		fmt.Printf("%d\n", i)
		i++
	}
}