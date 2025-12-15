package main

import "fmt"

func main() {

	/*
		Tabla ASCII
		65 a 90 Alfabeto mayúsculas
		97 a 122 Alfabeto minúsculas
	    salida: número que corresponde a la letra, la letra, usa for anidados. En mayúsculas y minúsculas
		97
		a
		a
		a
		a
		98
		b
		b
		b
		b
	*/

	fmt.Println("Mayúsculas")
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j <= 4; j++ {
			fmt.Printf("\t%c\n", i)
		}
	}

	fmt.Println("\nMinúsculas")
	for i := 97; i <= 122; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 4; j++ {
			fmt.Println("\t", string(i))
		}
	}
}
