// control statements
// packages
// declarations
// operators

package main

import (
	"fmt"
	"os"
)

func control_statements() {
	if a := 1; a == 1 {
		fmt.Println("a is one")
	} else {
		fmt.Println("a is not one")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("(%d %d)\n", i, i*i)
	}

	// only loop for-loop

	myArray := [...]int{1, 2, 3}
	// for i := range myArray {
	for i, v := range myArray { // here v get copy of myArray[i]
		fmt.Println(i, myArray[i])
		fmt.Println(i, v)
	}

	/*
		for k := range myMap {}
		for k, v := range myMap {}
	*/

	for {
		// infinite loop
		break // not now
	}

	/*
		outer:
			for ... {
				for ... {
					if ... {
						break / continue outer
					}
				}
			}
	*/

	switch a := 77; a {
	case 0, 1, 2:
		fmt.Println("underflow possible")
	case 3, 4, 5, 6, 7, 8:
		// empty
		fallthrough // by default no fallthrough means break
	default:
		fmt.Println("warning: overload")
	}

	// switch on true
	/*
		switch {
		case a <= 2:
			// st
		case a <= 8:
			// st
		default:
		}
	*/
}

func packages() {
	// used for information hiding / encaptulation
	// a package should embed deep functionality behind a simple API
	// best example is Unix file API

	// there are two type of scope
	// outside of function = package scope
	// inside of function = function scope
	// short declaration operator := only allowed in func scope

	// Every name that's capitalised is exported
	// that means another package in the program can import it
	// within a package, everything is visible even across files

	// each source file in your package must import what it needs
	// A package "A" cannnot import a package that imports A

	/*
		package pig

		const A = 1

		var B int = C
		var C int = A

		func init() {
			...
		}

		only the runtime can call init, also before main
		this init is private
		it is useful for loading plugins, driver for db, etc
		plugins is a little piece of code that you pull at runtime after main starts
	*/
}

func declarations() {
	/*
			const, type, var, :=, func,
			formal parameters and named returns of a function

		// this is called declaration block
		// this avoid multiple var keywords
		var (
			x, y int
			z float32
			s string
		)
	*/

	/*
		the short declaration operator := has some rules:
		1. It can't be used outside of a function
		2. It must be used (instead of var) in a control statement (if, etc)
		3. It must declare at least one new variable

		err := doSt()
		err := doStElse() // wrong bcz cannot be redeclared

		x, err := getSomeValue() // Ok; err is not redeclared

		4. It won't re-use an existing declaration from an outer scope

		shadowing is possible
	*/

	// BadRead(...) // shadowing example
}

func BadRead(f *os.File, buf []byte) error {
	var err error

	for {
		_, err := f.Read(buf) // shadows 'err' above

		if err != nil {
			break
		}
	}

	return err // will always be nil
}

func types() {
	/*
		STRUCTURAL TYPING

		it's the same type if it has the same structure or behaviour
		e.g. arrays of the same size and base type
		e.g. functions with the same parameter & return types
	*/

	/*
		NAMED TYPING
		Go uses named typing for non-function user-declared types
		type x int

		var a x  // x is a defined type; base int
		b := 2   // b default to int
		a = b    // type mismatch
		a = 12   // ok, untyped literal
		a = x(b) // ok, type conversion
	*/

	// Go keeps "arbitrary" precision for literal values (256 bits or more)
	// Interger literal are untyped
	// - assign a literal to any size integer without conversion
	// - assign an integer literal to float, complex also
}

func main06() {
	control_statements()
	fmt.Println()
	packages()
}
