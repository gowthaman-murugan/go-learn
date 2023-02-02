package main

import "fmt"

func main() {
	var a = "Hi"
	fmt.Println(a)

	var b, c int = 1, 2 // You can declare multiple variables at once.

	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e string = "Hello"
	fmt.Println(e)

	var f int // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	fmt.Println(f)

	g := 123 // := syntax is shorthand for declaring and initializing a variable
	fmt.Println(g)

	h := "apple" //  := syntax is shorthand for declaring and initializing a variable
	fmt.Println(h)
}
