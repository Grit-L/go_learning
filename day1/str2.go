package main

import "fmt"

// 单引号字符类型  双引号才是字符串
// len(str)
// fmt.Sprintf()  字符串拼接
// String.Contains
// String.Split 分割
// String.Index 
// String.LastIndex 
// String.Join
// byte类型 asscii
// rune类型 uint8
func main(){
	ch := '和'
	st := "hhh"
	f := `
	世界上只有一种真正的英雄主义，
	就是认清了生活的真相后还依然热爱它
	`
	fmt.Println(ch)
	fmt.Print(st)
	fmt.Printf(f)
	// 类型转换
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '白'  // rune(int32)
	fmt.Printf(string(s3))
}

