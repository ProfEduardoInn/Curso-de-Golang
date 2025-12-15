package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad uint8
	estatura float32
	sexo bool
}
type estudiante struct {
	persona  //también se puede p persona
	curso string
	grupo uint16
	nombre string
}

func main()  {

	estudiante1 := estudiante{
		persona: persona{ //también se puede p: persona{}
			nombre: "Lorenzo",
			apellido: "Lopez",
			edad: 19,
			estatura: 1.90,
		},
		curso: "Derecho",
		grupo: 4,
		nombre: "Lennon",
	}

	fmt.Println(estudiante1)
	fmt.Println(estudiante1.nombre)
	fmt.Println(estudiante1.persona.nombre)
	fmt.Printf("Hola, %s %s %s, tienes %d años de edad, mides %.02f de altura, eres estudiante de %s y estás en el grupo %d\n", estudiante1.persona.nombre, estudiante1.nombre, estudiante1.apellido, estudiante1.edad, estudiante1.estatura, estudiante1.curso, estudiante1.grupo)
}