package main

import (
	"fmt"
	"strings"
)

func main03() {
	s := "हिhello"

	fmt.Printf("%T %[1]v %v\n", s, len(s))
	fmt.Printf("%T %[1]v\n", []rune(s))
	fmt.Printf("%T %[1]v\n", []byte(s))

	s2 := "the quick brown fox"

	a := len(s2)
	b := s2[:3]
	c := s2[4:9]
	d := s2[:4] + "slow" + s2[9:] // O(n) bcz make a copy

	// s[5] = 'a' // Error
	s2 += "es" // O(n) bcz make a copy

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(s2)

	s2 = strings.ToUpper(s2)
	fmt.Println(s2)
}
