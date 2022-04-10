package main

import (
	"errors"
	"fmt"
)

func main() {
	errLearn()
	fmt.Printf("ending…… \n")

	err := errorDeal("sdf.nn")
	if err != nil {
		//输出错误,终止程序
		panic(err)
	}
	fmt.Printf("continue running")
}

func errLearn() {
	//错误捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("code error \n")
		} else {
			fmt.Printf("code runing \n")
		}

	}()
	n1 := 10
	n2 := 0
	n3 := n1 / n2
	fmt.Printf("hhh %v", n3)
}

func errorDeal(name string) (err error) {
	if name != " xx.ini" {
		//自定义错误
		return errors.New("文件名称错误……")
	}
	return nil
}
