package main

import "fmt"

func main() {
	//Operador AND y OR

	//Operador AND &&
	//AND necesita ser verdadera en ambas partes para que entre a la condicióm
	if edad := 10; edad > 18 && edad < 30 {
		fmt.Println("Eres joven AND")
		
	}
	//Operador OR ||
	//OR puede ser verdadera en una de las partes para que entre en la condición
	if edad := 10; edad > 18 || edad < 30 {
		fmt.Println("Eres joven OR")
	}

	if edad := 18; edad > 18 && edad < 30 {
		fmt.Println("Puedes pasar")
	} else if edad > 30 {
		fmt.Println("No puedes pasar, ya estás viejo.")
	} else if edad < 18 {
		fmt.Println("No puedes pasar, eres muy joven.")
	} else {
		fmt.Println("Ni idea, no sé que es esto.")
	}
}