package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main12() {
	var e = Employee{
		Name:   "Adarsh",
		Number: 100,
		Hired:  time.Now(),
	}

	b := Employee{
		Name:   "Abhijeet",
		Number: 101,
		Hired:  time.Date(2022, time.May, 15, 10, 30, 0, 0, time.Local),
	}

	e.Boss = &b

	fmt.Printf("%#v\n", e)
	fmt.Printf("%#v\n", b)

	c := map[string]Employee{}

	c["Abhijeet"] = e
	c["Adarsh"] = b

	// c["Abhijeet"].Boss = &c["Adarsh"]
	// both lhs and rhs are now allowed in map
	// like c[""].st and &c[""]; if c[""] is struct
	// Thats why always use *Struct

	/////////////////////////////////////////////////////////////////////////////////////////
	var a = struct {
		title string
	}{
		"The White Album",
	}

	var ptr_album *struct {
		title string
	}

	ptr_album = &a

	fmt.Println(ptr_album.title, (*ptr_album).title)
	// both are same

	/////////////////////////////////////////////////////////////////////////////////////////
	var a1 = album1{
		"The White Dog",
	}
	var a2 = album2{
		"The Black Dog",
	}

	// a1 = a2 // this is not allowed bcz they are diff name type
	a1 = album1(a2) // this possible bcz they are same structural type

	fmt.Println(a1, a2)

	// a set type (instead of bool)
	// var isPresent map[int]struct{}

	// a very cheap channel type
	// done := make(chan struct{})

	jsonEncoding()
}

type album1 struct {
	title string
}
type album2 struct {
	title string
}

// struct tags
// json use this and reflection concept to convert into json
type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func jsonEncoding() {
	r := Response{
		Page: 1,
		Words: []string{
			"up",
			"in",
			"out",
		},
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%+v\n", r)
	fmt.Println(string(j))
}
