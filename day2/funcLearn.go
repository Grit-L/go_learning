package main

import (
	"fmt"
	"github.com/trans"
)

func main() {
	fmt.Println("雪国列车")
	trans.Sd()
	//trans.Stest()
	a, b, c := operate(10, 2)
	fmt.Println(a, b, c)

	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)

	n := 0
	reply := &n
	Multiply(a, b, reply)
	fmt.Println(*reply, &reply)

	//匿名函数
	res := func(a1 int, b1 int) int {
		return a1 + b1
	}(1, 2)
	fmt.Printf("res: %d\n", res)

	//fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
}

func operate(a, b int) (ressum int, resc int, ressub int) {
	ressum = a + b
	resc = a * b
	ressub = a - b

	return
}
func operatev1(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a < b
		min = b
		max = a
	}
	return
}

//this function changes reply
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
