package main

import "fmt"

type carro struct{
	color, marca string
	modelo, numeroPuertas uint16
}

type camioneta struct {
	carro
	lujo, cuatroPuertas bool
}

type automovil struct {
	carro
	lujo, cuatroPuertas bool
}

func main()  {

/*
Crea un vehiculo con una estructura qye contenga datos como el color, marca, módelo y número de puertas.
- Genera dos tipos de vehículos como son camioneta y autómovil (carro si lo prefieres).
- Embebe el vehículo en ambos tipos de estructuras (camioneta y autómovil).
- Agrega a cada una de las estructuras (camioneta y autómovil) los campos de "lujo" y "cuatroPuertas" (como booleanos).
- Crea dos tipos de vehículos y asignales sus valores.
*/

vehiculo1 := camioneta{
	carro: carro{
		color: "blanco",
		marca: "Toyota",
		modelo: 2025,
		numeroPuertas: 4,
	},
		lujo: true,
		cuatroPuertas: true,

}

vehiculo2 := automovil{
	carro: carro{
		color: "Amarillo",
		marca: "Ford",
		modelo: 2015,
		numeroPuertas: 2,
	},
		lujo: false,
		cuatroPuertas: false,
}

fmt.Println(vehiculo1)
fmt.Println(vehiculo2)
	
}