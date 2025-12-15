package main

import "fmt"

func main() {

	//Dará erros de compilação
	var x uint8 = 255
	fmt.Printf("%v, %T\n", x, x)
	x++ // x = x+1
	//x-- x = x-1
	fmt.Printf("%v, %T\n", x, x)
	x += 10
	fmt.Printf("%v, %T\n", x, x)

	//Jeito certo
	var y uint16 = 324
	fmt.Printf("%v, %T\t", y, y)

	var z rune = 255
	fmt.Printf("%v, %T\n", z, z)

	var a int = 640
	fmt.Printf("%v, %T\n", a, a)
} 