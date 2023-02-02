package main

import "fmt"

func main() {
	fmt.Println("hello" + "world")
	fmt.Println("2+2", 2+2)
	fmt.Println("2-2", 2-2)
	fmt.Println("2*2", 2*2)
	fmt.Println("2/2", 2/2)

	fmt.Println("1.9+2.1", 1.9+2.1)
	fmt.Println("6.0/2.0", 6.0/2.1)

	fmt.Println("0=='0'", 0 == '0')
	fmt.Println("0==0", 0 == 0)
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!false", !false)

}
