package main

import (
	"fmt"
	"time"
)

const LIM = 100

var fibos [LIM]uint64

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
	num := 10
	fmt.Println("num addr:", &num)
	point(&num)
	fmt.Print(num)
	fmt.Println("\ncommn", comn(3))

	callback(10, Add)

	start := time.Now()
	//res := fibo(50)
	res := fiboCache(50)
	end := time.Now()
	fmt.Printf("运行时间：%s, 结果：%d", end.Sub(start), res)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func point(n1 *int) {
	*n1 = *n1 + 5
	fmt.Println("addrss: ", &n1)
	fmt.Println("*n1: ", *n1)
	fmt.Println(n1)
}

func comn(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * comn(n-1)
	}
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func fibo(n int) int {
	if n <= 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func fiboCache(n int) (res uint64) {
	if fibos[n] != 0 {
		res = fibos[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fiboCache(n-1) + fiboCache(n-2)
	}
	fibos[n] = res
	return
}
