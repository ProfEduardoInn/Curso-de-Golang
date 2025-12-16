package main

import "fmt"

type persona struct{
	nombre, apellidos []string
	edad uint8
	peso, estatura float32
	hobbies []string
}

func main()  {
	/*
	- Crea una estructura llamada persona con los siguientes campos:
	nombre, apellido, edad, peso, estatura y hobby.
	- Crea dos persona y muestra sus datos con un formato específico

	 Hola, mi nombre es: Juan, mi apellido es: Perez, tengo: 23 años, peso: 65kg, estatura 1.75 y mi hobby es: el futbol.
	*/

	persona1 := persona{
		nombre: []string{"Juan", "Carlos"},
		apellidos: []string{"Perez", "Lopez"},
		edad: 23,
		peso: 65.8,
		estatura: 1.75,
		hobbies: []string{"el Fútbol", "la danza"},
	}
	persona2 := persona{
		[]string{"Eduardo", "Luis"},
		[]string{"Rojas", "Ardila"},
		33,
		79.55,
		1.72,
		[]string{"cocinar", "programar"},
	}

	fmt.Printf("\nHola, mi nombre es: %s, mis apellidos sob: %s, tengo %d años, peso: %.02f, mido: %.02f, mis hobbies es: %s\n\n", persona1.nombre, persona1.apellidos, persona1.edad, persona1.peso, persona1.estatura, persona1.hobbies)
	
	fmt.Printf("Hola, mi nombre es: %s, mis apellidos son: %s, tengo %d años, peso: %.02f, mido: %.02f, mi hobbies son: %s\n", persona2.nombre, persona2.apellidos, persona2.edad, persona2.peso, persona2.estatura, persona2.hobbies)

	for i, v := range persona1.nombre {
		fmt.Println(i, v)
	}
	for i, v := range persona1.apellidos {
		fmt.Println(i, v)
	}
	for i, v := range persona1.hobbies {
		fmt.Println(i, v)
	}
	for i, v := range persona2.nombre {
		fmt.Println(i, v)
	}
	for i, v := range persona2.apellidos {
		fmt.Println(i, v)
	}
	for i, v := range persona2.hobbies {
		fmt.Println(i, v)
	}
}