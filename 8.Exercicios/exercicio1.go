package main

import "fmt"

/* Dicas
 - Use o operador de declaração curta.
 - Use o tipo de dado adequado para cada variável.
 - Use um formato específico para imprimir.
*/

/*Exercício 1
 - Crie 4 variáveis de tipos diferentes.
 - Asigne a uma delas tua idade.
 - a outra teu nome.
 - a outra se você está gostando do curso.
 - e por último o número pi com 2 decimais.
*/

func main() {
	edad := 33
	nombre := "Eduardo"
	gostou := true
	pi := 3.1416

/*Exercício 2
 - Imprima o resultado das variáveis na tela.
 - Cada um com seu tipo.
 - e sua representação en binário, octal e hexadecimal.
*/

	fmt.Println(edad, nombre, gostou, pi)
	fmt.Println()
	fmt.Printf("(tipo) T = %T, %T, %T, %T\n", edad, nombre, gostou, pi)
	fmt.Printf("(base 2 - binario) b = %b, %b, %b, %b\n", edad, nombre, gostou, pi)
	fmt.Printf("(base 8 - octal) o = %o, %o, %o, %o\n", edad, nombre, gostou, pi)
	fmt.Printf("(base 16 - Hexadecimal) hex = %x, %x, %x, %x\n", edad, nombre, gostou, pi)
}