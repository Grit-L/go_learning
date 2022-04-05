package main

import "fmt"

func main() {
	s := season(22)
	fmt.Println(s)
}

func season(x int) string {
	switch x {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Summer"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Spring"
	}
	return "unknow"
}
