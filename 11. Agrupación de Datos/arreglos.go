package main

import "fmt"

func main() {
	var anios [3]int
	fmt.Println(anios)
	anios[0] = 2026
	fmt.Println(anios)
	anios[1] = 2027
	fmt.Println(anios)
	anios[2] = 2028
	fmt.Println(anios)
	anios[0] = 3500
	fmt.Println(anios)
	
	hola := []string{"hola", "mundo", "!"}
	fmt.Println(hola)
	fmt.Println(len(hola))
	adios := []int{1,2,3,4,5,6,7,8,9,10,11,1215,134,165}
	fmt.Println(len(adios))
	var arreglo1 []string
	fmt.Println(len(arreglo1))
}