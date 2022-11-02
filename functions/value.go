package main

import "fmt"

func funcNum(a int) int {
	var x int
	x = a
	print(x)
	return x
}

func main() {
	var y int
	print("digite o valor: ")
	fmt.Scan(&y)
	funcNum(y)
}