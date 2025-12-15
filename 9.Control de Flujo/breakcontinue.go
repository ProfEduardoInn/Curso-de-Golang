package main

import "fmt"

func main() {

	//Break
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	//Continue
	for i := 0; i <= 6; i++ {
		fmt.Println(i)
		if i == 3 {
			continue
		}
		fmt.Println("Esto es una prueba")
	}
}