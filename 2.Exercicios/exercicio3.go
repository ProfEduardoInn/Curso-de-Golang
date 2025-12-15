package main

import "fmt"

/*Do exercício anterior asigne os valores de 50, "Capdesis", true (ao tipo de dado correspondente)

 - Use o String Format (Lembre-se do pacote fmt). Imprima todos os valores com apenas um print e na mesma linha.
 - Asigne eles a uma nova variável chamada resultado.
*/

//Variáveis globais
var idade int
var nome string
var like bool

func main() {

	idade = 50
	nome = "Capdesis"
	like = true
	resultado := fmt.Sprintf("\t%v, \t%v, \t%v\n", idade, nome, like)
	fmt.Println(resultado)
}