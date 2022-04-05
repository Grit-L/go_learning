package main

import (
	"fmt"
)

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func traceSim(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func aa() {
	defer un(traceSim("a"))
	fmt.Println("in a")
}

func bb() {
	defer un(traceSim("b"))
	fmt.Println("in b")
	aa()
}

func main() {
	b()
	fmt.Printf("<==================>\n")
	bb()

}
