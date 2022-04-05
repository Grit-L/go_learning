package main

var ac string

func main() {
	ac = "G"
	print(ac)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(ac)
}
