package main

import "fmt"


func main()  {
	coche := struct {
		marca string
		modelo string
		color string
		velocidad uint16
	} {
		marca: "Ford",
		modelo: "Fiesta",
		velocidad: 100,
		color: "Rojo",
	}
	fmt.Println(coche)
}