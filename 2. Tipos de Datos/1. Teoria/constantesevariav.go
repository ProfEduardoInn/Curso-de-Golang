package main

import "fmt"

func main() {

	const (
		nome      string  = "Eduardo"
		idade     uint8   = 33
		pi        float32 = 3.1415
		sobrenome string  = "Rojas"
	)
	var (
		anome      string  = "Eduardo"
		aidade     uint8   = 33
		api        float32 = 3.1415
		asobrenome string  = "Rojas"
	)

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(pi)
	fmt.Println(sobrenome)
	fmt.Printf("%T\n", nome)
	fmt.Printf("%T\n", idade)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", sobrenome)

	fmt.Println(anome)
	fmt.Println(aidade)
	fmt.Println(api)
	fmt.Println(asobrenome)
	fmt.Printf("%T\n", anome)
	fmt.Printf("%T\n", aidade)
	fmt.Printf("%T\n", api)
	fmt.Printf("%T\n", asobrenome)

	anome = "Luis"
	fmt.Println(anome)
}
