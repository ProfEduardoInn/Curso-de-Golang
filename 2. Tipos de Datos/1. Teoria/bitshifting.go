package main

import "fmt"

//Bit Shifting

 func main() {
	a := 5
	fmt.Printf("%T, %v, %b\n", a, a, a)
	b := a << 2
	fmt.Printf("%T, %v, %b\n", b, b, b)
	
	c := 5
	fmt.Printf("%T, %v, %b\n", c, c, c)
	d := c >> 2
	fmt.Printf("%T, %v, %b", d, d, d)
 }
