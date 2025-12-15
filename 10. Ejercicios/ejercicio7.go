package main

import "fmt"

func main() {

var edad uint8
	var nombre string
	fmt.Println("ingresa tu edad: ")
	fmt.Scanln(&edad)
	fmt.Println("ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Printf("Hola, %s, tu edad es: %d\n", nombre, edad)


	switch edad {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("Eres un bebé")
	case 6, 7, 8, 9, 10:
		fmt.Println("Eres un niño")
	case 11, 12, 13, 14, 15, 16, 17:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres mayor de edad")
	}

	switch {
	case edad < 18:
		fmt.Println("Eres menor de edad")
	case edad > 18 && edad < 30:
		fmt.Println("Eres un adulto joven")
	case edad > 30 && edad < 55:
		fmt.Println("Eres un adulto mayor")
	default:
		fmt.Println("Eres un anciano")
	}

	switch nombre { 
	case "Juan", "juan":
		fmt.Println("Tu nombre es Juan")
	case "Pedro", "pedro":
		fmt.Println("Tu nombre es Pedro")
	case "Eduardo", "eduardo":
		fmt.Println("Tu nombre es Eduardo")
	default:
		fmt.Println("Tu nombre no está registrado")
	 }
}

}