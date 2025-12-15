package main

import "fmt"

func main() {

	text := "\"hola\""
	textoDiferente := `
Hola
	`
	fmt.Println(text)
	fmt.Printf("%T\n", text)
	fmt.Println(len(text))
	fmt.Println("--------------------")

	fmt.Println(textoDiferente)
	fmt.Printf("%T\n", textoDiferente)
	fmt.Println(len(textoDiferente))

	//Slice de bytes
	sliceBytes := []byte(textoDiferente)
	fmt.Println(sliceBytes)
	fmt.Printf("%T\n", sliceBytes)

	fmt.Println(sliceBytes[0])
	fmt.Println(sliceBytes[3]) 
	fmt.Println(sliceBytes[6]) 
	fmt.Println(len(sliceBytes))
	fmt.Println(len(sliceBytes)-1)
	fmt.Println(sliceBytes[len(sliceBytes)-1])

	sliceBytes[0] = 108
	fmt.Println(sliceBytes)
	fmt.Println(len(sliceBytes))
}