package main

import "fmt"

func main() {

	mapa := map[string]int{
		"Pedro":   31945502918,
		"Lucas":   31947893312,
		"Gabriel": 31977864314,
		"Paula":   31988974654,
		"Ana":     38794113456,
	}

	maps := map[string]string{
		"Pedro":   "31945502918",
		"Lucas":   "31947893312",
		"Gabriel": "31977864314",
		"Paula":   "31988974654",
		"Ana":     "38794113456",
	}

	fmt.Println("Mapa 1", mapa)
	fmt.Println("Mapa 2", maps)
	fmt.Println("Mapa 1 Len", len(mapa))
	fmt.Println("Mapa 2 Len", len(maps))

	fmt.Println(mapa["Pedro"])
	fmt.Printf("Pedro Mapa 1 = %T\n", mapa["Pedro"])
	fmt.Println(maps["Gabriel"])
	fmt.Printf("Pedro Mapa 2 = %T", maps["Pedro"])
	fmt.Println(maps["Eduardo"])
	fmt.Printf("Eduardo Mapa 2 = %T\n", maps["Eduardo"])
	fmt.Println(mapa["Eduardo"])
	fmt.Printf("Eduardo Mapa 2 = %T\n", mapa["Eduardo"])


	//v, ok := mapa["Eduardo"]
	//fmt.Println(v, ok)

	_, ok := mapa["Eduardo"]
	fmt.Println(ok)

	/*
	if _, ok := mapa["Pedro"]; ok {
		fmt.Println("Existe")
	} else {
		fmt.Println("No existe")
	}
	*/
	if v, ok := mapa["Pedro"]; ok {
		fmt.Println("El número de Ana es: ", v)
	} else {
		fmt.Println("No existe")
	}

}