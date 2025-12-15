package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad uint8
	estatura float32
	sexo bool
}

func main()  {

	persona1 := persona{
		nombre: "Pedro",
		apellido: "Busquet",
		edad: 17,
		estatura: 1.72,
		sexo: true,
	}

	persona2 := persona{
		"Eduardo",
		"Rojas",
		33,
		1.72,
		true,
	}

	persona3 := persona{
		apellido: "Romero",
		edad: 21,
		sexo: false,
		nombre: "Lucas",
		estatura: 1.82,
	}
	
	/*persona4 := persona{
		"Romero",
		21,
		alse,
		"Lucas",
		1.82,
	}*/

	fmt.Println(persona1)
	fmt.Println(persona1.apellido)
	fmt.Println(persona1.nombre, persona1.apellido, persona1.edad, persona1.estatura, persona1.sexo)
	fmt.Printf("Te llamas %s %s, tienes %d años, mides %.02f\n", persona1.nombre, persona1.apellido, persona1.edad, persona1.estatura)
	fmt.Println(persona2)
	fmt.Println(persona3)
	
}