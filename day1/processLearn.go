package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	//n := 10
	//if n > 1 {
	//	fmt.Printf("'q'")
	//} else if n < 1 {
	//	fmt.Printf("'w'")
	//} else {
	//	fmt.Printf("'e'")
	//}

	//if n > 11 {
	//	return
	//}
	//return

	// for n:=0; n<10; n++{
	// 	fmt.Print(n)
	// }

	// s := "0h世间好物不坚牢，彩云易散琉璃脆"
	// for i, v  := range s {
	// 	// i是字符的index 1个汉字3个index
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	s := isGreater(2, 5)
	fmt.Println(s)

	//如何通过在初始化语句中获取函数值
	if value := isGreater(4, 1); value {
		fmt.Printf("this is true\n")
	} else {
		fmt.Printf("false!!!")
	}
	atoi("12")

	if an, err := mySqrt(-14.0); err {
		fmt.Println(an)
	} else {
		fmt.Println(err)
	}
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

//将字符串转换为整数时，且确定转换一定能够成功时，可以将 Atoi 函数进行一层忽略错误的封装
func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case
	return math.Sqrt(f), true
}
