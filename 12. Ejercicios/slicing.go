package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//salida 1: [2 3 4 5]
	//salida 2: [1 2 3 4 5 6 7 8 9 10]
	//salida 3: [6 7 8 9 10]
	//salida 4: [7 8 9]

	fmt.Println("Salida 1: ", slice[1:5])
	fmt.Println("Salida 2: ", slice[:])
	fmt.Println("Salida 3: ", slice[5:])
	fmt.Println("Salida 4: ", slice[6:9])

}
