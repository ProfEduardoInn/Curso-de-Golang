package main

import "fmr"

func main() {
	x := []int{1,2,3,4,5,6,7,8,9,10,
	11,12,13,14,15,16,17,18,19,20}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[:])
	fmt.Println(x[4:])
	fmt.Println(x[:3])
	fmt.Println(x[:5])
	fmt.Println(x[2:8])
	fmt.Println(x[:9])
}