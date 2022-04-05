package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "this is string"
	fmt.Println(strings.HasPrefix(str1, "this"))

	var str2 string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str2, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Burger"))

	oldStr := "hshh old old h"
	fmt.Printf(strings.Replace(oldStr, "old", "new", 1) + "\n")

	fmt.Println(strings.Count(oldStr, "h"))

	repeatStr := "Over "
	fmt.Println(strings.Repeat(repeatStr, 3))

	fmt.Println(strings.ToLower(repeatStr))
	fmt.Println(strings.ToUpper(repeatStr))

	spaceString := " hhs fdsf "
	fmt.Println(strings.TrimSpace(spaceString))
	fmt.Println(strings.Trim(spaceString, " "))
	fmt.Println(strings.TrimLeft(oldStr, "h"))

	splitString := "hs jjhsd  kfd"
	fmt.Println(strings.Fields(splitString))
	fmt.Println(strings.Split(splitString, "s"))

	sliceStr := strings.Split(splitString, "s")
	joinStr := strings.Join(sliceStr, ":")
	fmt.Println(strings.Replace(joinStr, " ", "", -1))
}
