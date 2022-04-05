package main

import (
	"fmt"
	"strings"
)

func main() {
	// 九九乘法表
	//for x := 0; x < 9; x++ {
	//	// fmt.Println(x)
	//	for i := 1; i <= x; i++ {
	//		num := i * x
	//		fmt.Printf("%d*%d=%d  ", x, i, num)
	//	}
	//	fmt.Println("\n")
	//}

	//range15()
	//fizzBuzz()
	//str2 := "Chinese: 优美的中国话"
	//forRange(str2, "rune")

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
		fmt.Printf("%d ", v)
	}

	//无限循环
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

	//s := ""
	//for s != "aaaaa" {
	//	fmt.Println("Value of s:", s)
	//	s = s + "a"
	//}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	//	i := 0
	//START:
	//	fmt.Printf("count:%d ", i)
	//	i++
	//	if i < 15 {
	//		goto START
	//	}
}

func range15() {
	for i := 1; i < 26; i++ {
		//fmt.Println(i)
		fmt.Printf(strings.Repeat("G", i) + "\n")
	}
}

func fizzBuzz() {
	for i := 1; i < 101; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func forRange(str string, t string) {
	if t != "rune" {
		for pos, char := range str {
			fmt.Printf("Character on position %d is: %c \n", pos, char)
		}
	} else {
		for index, rune := range str {
			fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
		}
	}
}
