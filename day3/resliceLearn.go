package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice1 is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(len(a[1:1]))
	a = a[0:2]
	for _, v := range a {
		fmt.Printf("%d \n ", v)
	}

	//	复制与追加
	slFrom := []int{1, 2, 3, 4}
	slTo := make([]int, 5)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	/*
		 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素是同类型。
		如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。
		因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了
	*/
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3[:1], 4, 5, 6)
	fmt.Println(sl3)

	target := []string{"w", "x", "y"}
	source := []string{"a", "b", "c", "d", "e"}
	insertStringSlice(2, target, source)

	removeStringSlice(2, 4, source)
}

//写一个函数 insertStringSlice 将切片插入到另一个切片的指定位置
func insertStringSlice(i int, newSlice []string, source []string) (r []string) {
	r = append(r, source[:i]...) // slice后加...解包
	r = append(r, newSlice[:]...)
	r = append(r, source[i:]...)
	fmt.Println(r)
	return
}

//写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片中移除
func removeStringSlice(start int, end int, source []string) (res []string) {
	for i := 0; i < len(source); i++ {
		if i >= start && i <= end {
			continue
		}
		res = append(res, source[i])
	}
	fmt.Println(res)
	return
}
