package main

import "fmt"

func main() {

	const (
		ano1 = 1992 + iota
		ano2
		ano3
		ano4
		ano5
	)
	fmt.Println(ano1)
	fmt.Println(ano2)
	fmt.Println(ano3)
	fmt.Println(ano4)
	fmt.Println(ano5)
	
	fmt.Println("\nEjercicios con for")
	
	anios := 0
	for i := 1992; i <= 2025; i++ {
		fmt.Printf("\nAño %v: %d\n", anios, i)
		anios++
	}
	fmt.Printf("\nPasaron %v años", anios - 1)
}