package main

import "fmt"

/* Declare 3 variáveis: 1 de tipo inteiro, 1 de tipo string e 1 de tipo bool.
Com os identificadores de idade, nome e "se deixou seu like" */

/* Considerações
 - As variáveis tem que ser globais.
 - Não pode usar o operador de declaração curta.
 - Não pode declarar e inicializar as variáveis na mesma linha.
 - Imprima o valor asignado a cada variável sem asfgnar nenhum valor, depois asigne os valores e imprima eles de novo.
 - Como é nomeado o valor que é asignado para as variáveis ao ser declaradas (sem ser inicializadas)
*/

//Variáveis globais
var idade int
var nome string
var like bool

func main() {

	fmt.Println("Variáveis sem valor")
	fmt.Printf("%v, %v, %v\n", idade, nome, like)

	fmt.Println("\nVariáveis com valor")
	idade = 33
	nome = "Eduardo"
	like = true
	fmt.Printf("%v, %v, %v\n", idade, nome, like)

	fmt.Println("\nChama-se \"Zero\" value")
}