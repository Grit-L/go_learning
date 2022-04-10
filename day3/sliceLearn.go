package main

import (
	"fmt"
)

func main() {
	//	定义
	var array01 = [6]int{1, 7, 9, 0, 3, 8}
	var slice01 = array01[:]
	var slice02 = &array01
	slice03 := []int{2, 4, 8}
	slice04 := array01[:4]

	for _, v := range slice01 {
		fmt.Printf("%v \n", v)
	}
	slice02[2] = 99
	fmt.Printf("array01[2]=%d\n", array01[2])
	for _, v := range slice02 {
		fmt.Printf("%v \n", v)
	}
	for _, v := range slice03 {
		fmt.Printf("%v \n", v)
	}
	//切片本身就是一个指针
	fmt.Printf("%p \n", &array01)
	fmt.Printf("%p \n", slice02)
	fmt.Printf("%p \n", &slice02)

	//长度，容量
	fmt.Printf("slice04 cap is: %d, len is %d \n", cap(slice04), len(slice04))
	fmt.Printf("array len is:%d, slice len is:%d", cap(array01), len(slice02))

	csum := sum(slice03)
	fmt.Printf("%d", csum)

	//new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 *T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
	//make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channe
	slice1 := make([]int, 10)
	slice2 := make([]int, 5, 10)
	slice3 := new([5]int)[0:3]
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	fmt.Printf("The capacity of slice2 is %d, length is %d\n", cap(slice2), len(slice2))
	fmt.Printf("The capacity of slice3 is %d, length is %d\n\n", cap(slice3), len(slice3))

	s := make([]byte, 5)
	fmt.Printf("The capacity of s is %d, length is %d\n", cap(s), len(s))
	s = s[2:4]
	fmt.Printf("The capacity of s is %d, length is %d\n", cap(s), len(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]

	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c \n", s2[i])
	}
	s2[1] = 't'
	fmt.Printf("s1[1] =%c, s2[1] =%c \n", s1[1], s2[1])

	av := averageNum(slice03)
	fmt.Println(av)

	minS := minSlice(slice04)
	fmt.Printf("min value is %d \n", minS)
	maxS := maxSlice(slice04)
	fmt.Printf("max value is %d", maxS)

	//string的处理
	s4 := "dsjakjfkskfjk"
	//类型转换
	slice05 := []byte(s4)
	//处理中文
	slice06 := []rune(s4)
	fmt.Println(slice05, slice06)

	/*
		如何理解 new、make、slice、map、channel 的关系?
		1.slice、map 以及 channel 都是 golang 内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而 make 就是对三者进行初始化的一种操作方式
		2. new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以 slice、map、channel 需要 make 进行初始化并获取对应的内存地址，而非 new 简单的获取内存地址
	*/
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func averageNum(a []int) float32 {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return float32(sum / len(a))
}

func minSlice(s []int) int {
	min := s[0]
	for i := 1; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}
	return min
}

func maxSlice(s []int) int {
	min := s[0]
	for _, v := range s {
		if min < v {
			min = v
		}
	}
	return min
}
