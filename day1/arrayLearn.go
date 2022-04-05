package main

import "fmt"

// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
var doubleArray = [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

// 上面的声明可以简化，直接忽略内部的类型
var easyArray = [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

func array_test() {

	fmt.Println(doubleArray[0][1])
}

func main() {
	array_test()
}
