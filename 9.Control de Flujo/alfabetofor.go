package main

import "fmt"

func main() {
/*	//imprimir el abecedario Mayúsculas
	for i := 65; i <= 90 ; i++ {	
		fmt.Printf("%c ", i)
	}
	//Imprimir el abecedario Minúsculas
	for i := 97; i <= 122 ; i++ {	
		fmt.Printf("%c ", i)
	}
	//Con for tipo While
	j := 65
	for j <= 90 {
		fmt.Printf("%c ", j)
		j++
	}
	x := 97
	for x <= 122 {
		fmt.Printf("%c ", x)
		x++
	}
	//imprimir el abecedario Mayúsculas
	for i := 90; i >= 65 ; i-- {	
		fmt.Printf("%c ", i)
	}
	//Imprimir el abecedario Minúsculas
	for i := 122; i >= 97 ; i-- {	
		fmt.Printf("%c ", i)
	}
	//Con for tipo While
	z := 90
	for z >= 65 {
		fmt.Printf("%c ", z)
		z--
	}
	a := 122
	for a >= 97 {
		fmt.Printf("%c ", a)
		a--
	}
*/
	//Abecedario sin vocales
	//En mayúsculas
	for i := 65; i <= 90; i++ {
		if i == 65 || i == 69 || i == 73 || i == 79 || i == 85 {
			continue
		}
		fmt.Printf("%c ", i)
	}
	//En minúsculas
	x := 97
	for x <= 122 {
		if x == 97 || x == 101 || x == 105 || x == 111 || x == 117 {
			x++
			continue
		}
		fmt.Printf("%c ", x)
		x++
	}
}