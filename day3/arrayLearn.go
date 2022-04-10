package main

import (
	"fmt"
)

func main() {
	var arrTest [3]int
	//打印地址
	fmt.Printf("数组地址：%p,第一个：值地址 %p, %p \n", &arrTest, &arrTest[0], &arrTest[1])

	//初始化数组
	var array01 [3]int = [3]int{1, 2, 3}
	var array02 = [3]int{4, 5, 6}
	var array03 = [...]int{7, 8, 9}
	var array04 = [3]int{0: 44, 2: 33, 1: 22}
	fmt.Printf("array01: %v, array02: %v,array03: %v,array04: %v \n", array01, array02, array03, array04)

	//遍历
	for i := 0; i < len(array04); i++ {
		fmt.Println(arrTest[i])
	}

	for i, v := range array03 {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	items := [...]int{10, 20, 30, 40, 50}
	s := items[:]
	for i, item := range s {
		s[i] = item * 2
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
