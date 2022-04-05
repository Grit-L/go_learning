package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	a = iota
	b = iota
	c // 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值
	d = 66
	e
)

// 使用 iota 结合 位运算 表示资源状态的使用案例
const (
	Open    = 1 << iota // 0001
	Close               // 0010
	Pending             // 0100
)

const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)

var (
	x int
	y bool
	z string
)

func main() {
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	fmt.Println("hello world!!!")
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, z)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
