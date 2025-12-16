package main

import "fmt"

/*
 - Do exercício anterior, crie uma variável global sem o operardor de asginação curto, com o identificador idadePessoa com um tipo de dado que venha desde anostidos. 
 IMPORTANTE
(Não dá, golang não permite fazer direto assim, para fazer casting esta variável tem que ter o valor ao qual será convertido e não valor desde o qual será convertido)
 - Use a conversão para passar o valor guardado em anostidos com o tipo de dado que ele veio, do que derivou.
 - Use o operador de declaração curta para asignar-lhe o valor a idadePessoa.
 - Copie e cole o código do exercício 4
*/

type idade uint8
var anostidos idade
var idadePessoa uint8

func main() {
	fmt.Printf("O valor é: %v e o tipo de dado é: %T\n", anostidos, anostidos)
	anostidos = 20
	fmt.Printf("O novo valor é: %v", anostidos)
	idadePessoa = uint8(anostidos)
	fmt.Printf("\nO novo valor é: %v, %T\n", idadePessoa, idadePessoa)
	idadePessoa = 18
	fmt.Printf("O novo valor é: %v, %T\n", idadePessoa, idadePessoa)
}