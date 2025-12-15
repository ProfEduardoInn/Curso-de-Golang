package main

import "fmt"

func main() {
	var saludo string = "Hola"
	var conversa string = `
	"(Saludos 
	a todos,
	espero que la estén pasando
	muy bien con el curso de GO").
	`
	
	fmt.Printf("La variable saludo tiene un string: %s, Su tipo de dato es: %T", saludo, saludo)
	fmt.Printf("La variable conversa tiene un string: %s, Su tipo de dato es: %T", conversa, conversa)
}