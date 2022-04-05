package main

import (
	"fmt"
	"log"
	"strings"
)

var where = log.Print

func main() {
	fmt.Println(f())
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))

	where()
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	var f = addUpper()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	where()
	bmp := MakeAddSuffix("bmp")
	filename := bmp("ss.")
	fmt.Printf("\n%s", filename)
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

//闭包
//闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
func addUpper() func(int) int {
	var x int
	return func(b int) int {
		x += b
		return x
	}
}

//工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

//2
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
