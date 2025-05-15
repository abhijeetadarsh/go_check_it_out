package main

import (
	"encoding/json"
	"fmt"
)

func main10() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("%d %d %T %5t %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(v), cap(v), v, v == nil)

	/*
		s: ┌───┬───┬───┐
		   │ 0 │ 0 │ 0 │
		   └───┴───┴───┘

		t: ┌───┬───┬──────┐
		   │ 0 │ 0 │ 0x11 │─────▶ struct{}
		   └───┴───┴──────┘
		u: ┌───┬───┬──────┐
		   │ 5 │ 5 │ 0xbc │─────▶┌───┬───┬───┬───┬───┐
		   └───┴───┴──────┘      │ 0 │ 0 │ 0 │ 0 │ 0 │
		                         └───┴───┴───┴───┴───┘

		v: ┌───┬───┬──────┐
		   │ 0 │ 5 │ 0xcd │─────▶┌───┬───┬───┬───┬───┐
		   └───┴───┴──────┘      │   │   │   │   │   │
		                         └───┴───┴───┴───┴───┘
	*/

	// nil and empty behave diff
	j1, _ := json.Marshal(s)
	fmt.Println(string(j1)) // null

	j2, _ := json.Marshal(t)
	fmt.Println(string(j2)) // []

	// thus always check for if len(s) == 0

	// you can append to a nil slice
	// it is not that same as inserting into a nil map

	a := [3]int{1, 2, 3}
	b := a[:1] // only control the len not cap
	c := b[:2] // WTF

	fmt.Println("a =", a, len(a), cap(a))
	fmt.Println("b =", b, len(b), cap(b))
	fmt.Println("c =", c, len(c), cap(c))

	d := a[0:1:1] // control both len and cap
	// len j - i and cap k - i
	fmt.Println("d =", d, len(d), cap(d))

	/*
		a := [...]int{1, 2, 3}
		b := a[0:1]
		c := b[0:2:2]    // WTF

		fmt.Printf("a[%p] = %v\n", &a, a)
		fmt.Printf("b[%p] = %#[1]v\n", b)
		fmt.Printf("c[%p] = %#[1]v\n", c)

		c[0] = 9
		fmt.Printf("a[%p] = %v\n", &a, a)
		fmt.Printf("c[%p] = %#[1]v\n", c)

		c = append(c, 5)
		fmt.Printf("a[%p] = %v\n", &a, a)
		fmt.Printf("c[%p] = %#[1]v\n", c)
	*/
	// the above show that slice is basically a view of an underlying array
}
