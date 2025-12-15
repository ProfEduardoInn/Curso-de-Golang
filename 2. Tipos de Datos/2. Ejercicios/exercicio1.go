package main

import "fmt"

/*
Escreva um programa que imprima um nímero inteiro, decimal, binário e hexadecimal.
 - Se possível peça os dados ao usuário.
*/

func main() {

	var (
		inteiro     int64
		decimal     int64
		binario     int64
		hexadecimal int64
		idade       uint8
		nome        string
	) 

	fmt.Println("Escreva um número que deseja que seja inteiro: ")
	fmt.Scan(&inteiro)
	fmt.Println("Escreva um número que deseja que seja decimal: ")
	fmt.Scan(&decimal)
	fmt.Println("Escreva um número que deseja que seja binário: ")
	fmt.Scan(&binario)
	fmt.Println("Escreva um número que deseja que seja hexadecimal: ")
	fmt.Scan(&hexadecimal)
	fmt.Println("Escreve tua idade: ")
	fmt.Scan(&idade)
	fmt.Println("Escreve teu nome: ")
	fmt.Scan(&nome)

	fmt.Printf("%v\n", inteiro)
	fmt.Printf("%d\n", decimal)
	fmt.Printf("%b\n", binario)
	fmt.Printf("%x\n", hexadecimal)
	fmt.Printf("Olá, %s!, %v, %d, %b, %x", nome, idade, idade, idade, idade)
}