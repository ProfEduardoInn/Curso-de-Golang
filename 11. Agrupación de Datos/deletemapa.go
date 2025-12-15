package main

import "fmt"

func main() {
	mapa := map[string]int {
		"uno": 1,
		"dos": 2,
		"tres": 3,
	}
	fmt.Println(mapa)
	mapa["cuatro"] = 4
	fmt.Println(mapa)

	//Delete: delete(nameMap, key)
	delete(mapa, "cuatro")
	delete(mapa, "dos")
	fmt.Println(mapa)
	
	if valor, ok := mapa["cuatro"]; ok {
		fmt.Println(valor)
		delete(mapa, "cuatro")
	} else {
		fmt.Println("No existe la clave 4")
	}
}