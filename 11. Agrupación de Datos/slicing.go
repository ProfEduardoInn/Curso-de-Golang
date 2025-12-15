package main

import "fmt"

func main() {

	x := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("x = ", x)
	x = append(x, 11)
	fmt.Println("x = ", x)
	y := append(x, 12)
	fmt.Println("y = ", y)
	fmt.Println("x = ", x)
	x = append(x,y...)
	fmt.Println("x = ", x)
}