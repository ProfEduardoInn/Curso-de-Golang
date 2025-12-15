package main

import "fmt"

func main() {
	slice := make([]int, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	slice[0] = 1
	slice[2] = 12
	fmt.Println(slice)
	slice[9] = 13
	fmt.Println(slice)
	slice = append(slice, 14)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	slice = append(slice, 24)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}