package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota + 1
	e
	f
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Printf("%T", a)
}