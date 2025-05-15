package main

import "fmt"

func main03() {
	a, b := 2, 20.1
	var c complex128
	var (
		d bool
		e = "dog"
	)
	var f error

	fmt.Printf("a: %T %[1]v\n", a)
	fmt.Printf("b: %T %v\n", b, b)
	fmt.Printf("c: %T %v\n", c, c)
	fmt.Printf("d: %T %v\n", d, d)
	fmt.Printf("e: %T %v\n", e, e)
	fmt.Printf("f: %T %v\n", f, f)

	a = int(b)
	b = float64(a)

	fmt.Printf("a: %T %[1]v\n", a)
	fmt.Printf("b: %T %v\n", b, b)

	const (
		g int    = 1
		h uint   = 2
		i string = "a string"
		j bool   = true
	)
}
