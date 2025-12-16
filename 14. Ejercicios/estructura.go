package main

import "fmt"

type persona struct{
	nombre, apellido string
	edad uint8
	peso, estatura float32
	hobby string
}

func main()  {
	/*
	- Crea una estructura llamada persona con los siguientes campos:
	nombre, apellido, edad, peso, estatura y hobby.
	- Crea dos persona y muestra sus datos con un formato específico

	 Hola, mi nombre es: Juan, mi apellido es: Perez, tengo: 23 años, peso: 65kg, estatura 1.75 y mi hobby es: el futbol.
	*/

	persona1 := persona{
		nombre: "Juan",
		apellido: "Peraz",
		edad: 23,
		peso: 65.8,
		estatura: 1.75,
		hobby: "el Fútbol",
	}
	persona2 := persona{
		"Eduardo",
		"Rojas",
		33,
		79.55,
		1.72,
		"cocinar",
	}

	fmt.Printf("Hola, mi nombre es: %s, mi apellido es: %s, tengo %d años, peso: %.02f, mido: %.02f, mi hobby es: %s\n", persona1.nombre, persona1.apellido, persona1.edad, persona1.peso, persona1.estatura, persona1.hobby)
	
	fmt.Printf("Hola, mi nombre es: %s, mi apellido es: %s, tengo %d años, peso: %.02f, mido: %.02f, mi hobby es: %s\n", persona2.nombre, persona2.apellido, persona2.edad, persona2.peso, persona2.estatura, persona2.hobby)
}