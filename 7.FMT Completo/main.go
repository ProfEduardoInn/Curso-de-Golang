package main

import "fmt"

func main() {
	var edad = 20
	var booleano = true
	var flotante = 12.5
	var complejo = complex(12.5, 12.5)
	fmt.Println(edad)
	fmt.Printf("(tipo) T = %T\n", edad)
	fmt.Printf("(base 2 - binario) b = %b\n", edad)
	fmt.Printf("(valor) v = %v\n", edad)
	fmt.Printf("(valor GO) #v = %#v\n", edad)
	fmt.Printf("(base 10 - decimal) d = %d\n", edad)
	fmt.Printf("(booleano) t = %t\n", booleano)
	fmt.Printf("(caracter) c = %c\n", edad)

	//FLOTANTES
	fmt.Println()
	fmt.Printf("(flotante) %%b = %b\n", flotante)
	fmt.Printf("(flotante) %%e = %e\n", flotante)
	fmt.Printf("(flotante) %%E = %E\n", flotante)
	fmt.Printf("(flotante) %%f = %f\n", flotante)
	fmt.Printf("(flotante) %%F = %F\n", flotante)
	fmt.Printf("(flotante) %%g = %g\n", flotante)
	fmt.Printf("(flotante) %%G = %G\n", flotante)
	fmt.Printf("(flotante) %%x = %x\n", flotante)
	fmt.Printf("(flotante) %%X = %X\n", flotante)
	
	//COMPLEJOS
	fmt.Println("\nNúmeros Complejos")
	fmt.Printf("(complejos) %%b = %b\n", complejo)
	fmt.Printf("(complejos) %%e = %e\n", complejo)
	fmt.Printf("(complejos) %%E = %E\n", complejo)
	fmt.Printf("(complejos) %%f = %f\n", complejo)
	fmt.Printf("(complejos) %%F = %F\n", complejo)
	fmt.Printf("(complejos) %%g = %g\n", complejo)
	fmt.Printf("(complejos) %%G = %G\n", complejo)
	fmt.Printf("(complejos) %%x = %x\n", complejo)
	fmt.Printf("(complejos) %%X = %X\n", complejo)
}