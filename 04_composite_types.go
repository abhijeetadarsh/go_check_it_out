package main

import "fmt"

func arrays() { // arrays are passed by value; thus elements are copied
	var a [3]int
	var b = [3]int{1, 2, 3}
	var c = [...]int32{4, 5, 6} // sized by initializer

	var d [3]int
	d = b // element copied

	var m = [...]int{1, 2, 3, 4}
	// c = m // Error type mismatch

	fmt.Printf("%T: %[1]v\n%T: %[2]v\n%T: %[3]v\n%T: %[4]v\n%T: %[5]v\n", a, b, c, d, m)

	fmt.Printf("%T: %[1]v\n", m[1:3])
}

func slices() { // slices are passed by reference; no copying, updating OK
	var a []int
	var b = []int32{1, 2, 3}
	var c = []int{4, 5, 6} // sized by initializer

	a = append(a, 55)
	b = append(b, 56)

	d := make([]int, 4)
	// d = b // Error bcz they are slice of diff types

	var m = []int{1, 2, 3, 4}
	c = m // No copy of values

	fmt.Printf("%T: %[1]v\n%T: %[2]v\n%T: %[3]v\n%T: %[4]v\n%T: %[5]v\n", a, b, c, d, m)

	fmt.Printf("%T: %[1]v\n", m[1:3])
}

func maps() { // unordered map based on hash table
	var m map[string]int // nil, no storage
	p := make(map[string]int)

	// m["the"]++ // panic: assignment to entry in nil map
	p["the"]++

	fmt.Printf("%T: %[1]v\n", m)
	fmt.Printf("%T: %[1]v\n", p)

	var m2 = map[string]int{
		"and": 0,
		"the": 2,
		"or":  4,
	}

	// b := m == m2 // invalid operation: m == m2 (map can only be compared to nil)
	c := m == nil
	d := len(m2)
	// e := cap(m2) // invalid argument: m2 (variable of type map[string]int) for built-in cap

	fmt.Println(c, d, m2)

	f, ok := m2["and"]
	fmt.Println(f, ok)
	g, ok := m2["pig"]
	fmt.Println(g, ok)

	if _, ok := m2["the"]; ok {
		fmt.Print("we know _ is not the default value\n")
	}
}

func main04() {
	arrays()
	fmt.Println()
	slices()
	fmt.Println()
	maps()
}
