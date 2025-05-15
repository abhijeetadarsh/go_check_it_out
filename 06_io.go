package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func cat() {
	for _, fname := range os.Args[1:] {
		// fmt.Println(fname)
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		file.Close()
	}
}

func calc_size() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("The file:", fname, "is of size", len(data), "bytes")
		file.Close()
	}
}

func wc() {
	// word count

	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s) + 1
			lc++
		}

		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}

func main06() {
	calc_size()
	wc()
	return

	cat()
	a, b := 12, 245
	c, d := 1.2, 2.45

	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%X %x\n", a, b)
	fmt.Printf("%#x %#x\n", a, b)
	fmt.Printf("%f %.3f\n", c, d)

	fmt.Println()
	fmt.Printf("|%6d|%6d|\n", a, b) // here 6 means use atleast 6 columns
	fmt.Printf("|%06d|%06d|\n", a, b)
	fmt.Printf("|%-6d|%-6d|\n", a, b)

	fmt.Println()
	s := []int{1, 2, 3}
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	fmt.Println()
	e := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T\n", e)
	fmt.Printf("%q\n", e)
	fmt.Printf("%v\n", e)
	fmt.Printf("%#v\n", e)

	fmt.Println()
	m := map[string]int{"and": 1, "or": 2}
	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
}
