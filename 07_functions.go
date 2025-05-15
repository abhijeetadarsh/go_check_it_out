package main

func main09() {
	/*
		Functions are "first class" objects; you can:
		- Define them - even inside another function
		- Create anonymous function literals
		- Pass them as function parameters / return values
		- Store them in variables
		- Store them in slices and maps (but not as keys)
		- Store them as fields of a structure type
		- Send and recieve them in channels
		- Write methods against a function type
		- Compare a function var against nil
	*/

	// Function use structural typing
	/*
		All parameters are passed by copying something (i.e., by value)
		If the thing copied is a pointer or descriptor, then the shared
		backing store (array, hash table, etc.) can be changed through it

		Thus we think of it as "by reference"

		By value:
		- numbers
		- bool
		- arrays
		- structs

		By reference:
		- things passed by pointer (&x)
		- strings (but they're immutable)
		- slices
		- maps
		- channels
	*/

	Defer()
}

func Defer() {
	// The defer statement captures a function call to run later

	/*
		e.g. 1
		func main() {

			f, err := os.Open("my_file.txt")

			if err != nil {
				...
			}

			defer f.Close()
		}

		e.g. 2
		Notice that the defer will not execute when we leave the if block
		It only executes when func returns


		e.g. 3
		Unlike a closure, defer copies arguments to the deferred call

		func main() {
			a := 10
			defer fmt.Println(a)
			a = 11
			fmt.Println(a)
		}
		// prints 11, 10
		The parameter a gets copied at the defer statement (not a reference)

		e.g. 4
		func doIt() (a int) {
			defer func() {
				a = 2
			}()

			a = 1
			return
		}

		// returns 2
		We have a named return value and a "naked" return
		The deferred anonymous function can update that variable
		bcz named return value is same as local variable
	*/
}
