package main

import "fmt"

func main()  {
	/*
	- Crea una estructura anónima de una persona con por lo menos 4 atributos.
	*/

	persona := struct {
		nombre string
		apellido string
		edad uint32
		altura float32
		peso float32
		sexoM bool
		contactos map[string]int
		comidaFavorita []string
	} {
		nombre: "Pedro",
		apellido: "Busquet",
		edad: 17,
		altura: 1.68,
		peso: 72.2,
		sexoM: true,
		contactos: map[string]int{
			"Carlos": 2345897,
			"Sergio": 3455789,
		},
		comidaFavorita: []string{
			"Picanha",
			"Sushi",
			"Pollo frito",
		},
	}
	fmt.Println(persona)
	fmt.Println(persona.contactos["Carlos"])

	for i, v := range persona.comidaFavorita {
		fmt.Println(i, v)
	}
}