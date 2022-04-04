package main

import "fmt"

func main() {
	// n := 10
	// if n > 1 {
	// 	fmt.Printf("'q'")
	// } else if n < 1 {
	// 	fmt.Printf("'w'")
	// } else {
	// 	fmt.Printf("'e'")
	// }

	// for n:=0; n<10; n++{
	// 	fmt.Print(n)
	// }

	// s := "0h世间好物不坚牢，彩云易散琉璃脆"
	// for i, v  := range s {
	// 	// i是字符的index 1个汉字3个index
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	// 九九乘法表
	x := 0
	for x<9{
		x++
		// fmt.Println(x)
		for i:=1;i<=x;i++ {
			num := i*x
			fmt.Printf("%d*%d=%d  ",x,i,num)
		}
		fmt.Println("\n")
	}
}
