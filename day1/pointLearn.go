package main

import (
	"fmt"
)

func main() {
	var i3 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i3, &i3)

	var i1 = 6
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	fmt.Println(*(&i1))

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)

	//var p1 *int = nil
	//*p1 = 0
	right := 10
	left := 15
	swap(&right, &left)
	fmt.Println("right: ", right, "left: ", left)
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
