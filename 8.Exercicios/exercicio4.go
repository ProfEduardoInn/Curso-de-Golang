package main

import "fmt"

/* 
 - Crie seu própio tipo de dado que venha de um inteiro para a idade (a mais ótima possível)
 - Crie uma variável com seu novo tipo de identificador "idade" sem usar o operador de declaração curta.
 - Imprima o valor da variável idade, o tipo e asigne a ele o valor 20.
 - Imprima o valor de idade
*/

type idade uint8
var anostidos idade

func main() {

	fmt.Printf("O valor é: %v e o tipo de dado é: %T\n", anostidos, anostidos)
	anostidos = 20
	fmt.Printf("O novo valor é: %v", anostidos)
}