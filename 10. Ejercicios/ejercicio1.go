package main

import "fmt"

func main() {
	//Imprtme los números del 1 al 100 con for init

	for i := 0; i <= 100 ; i++ {
		fmt.Printf("%d\n", i)
	}
    
	i := 101
	for i <= 200 {
		fmt.Printf("%d\n", i)
		i++
	}
}