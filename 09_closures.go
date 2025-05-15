package main

import (
	"fmt"
	"sort"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b // note this vars are reference
		return b
	}
}

func main09() {
	// escape analysis
	/*
		func doIt() *int {
			var b int
			...

			return &b
		}
		.           ^
		By seeing this the escape analysis of the
		compiler but the "b" in heap bcz it have to
		live long
	*/

	// Closure is runtime thing
	/*
		Look at func fib
		Here closure is what that is returned by fib
		closure = function + env variables
		f := fib   // here f is a func pointer
		f := fib() // here f is a closure

		The inner function gets a reference to the outer function's vars
		Those variable may end up with a much longer lifetime than expected
		as long as there's a reference to the inner function
	*/
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Print(x, " ")
	}
	fmt.Println()
	g := fib()
	for x := g(); x < 100; x = g() {
		fmt.Print(x, " ")
	}
	fmt.Println()
	// this is bcz f and g get different a and b
	// both have the same func pointer but the env pointers are different

	// another e.g.
	type kv struct {
		key string
		val int
	}
	var ss []kv
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})
	// this sort.Slice func wants a particular closure signature
	// look ss inside that func

	// below is a example that show env var are references
	// so that vars can mutate
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		// i := i // closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}

	/*
		Simple function
		fp ─────▶  Code:
				mov r0, #10
				mov r1, #3
				add r0, r0, r1
				...

		Closure
		+---------+
		|  env    | ─────▶  Environment:
		|         |         &a
		|  fp     |         &b
		+---------+
			│
			▼
		Code:
		mov r0, #10
		mov r1, #3
		add r0, r0, r1
		...
	*/
}
