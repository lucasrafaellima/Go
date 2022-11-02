package main

import "fmt"

func xpto(a*int) int {
	*a = 100
	return *a
}

func main() {
	b := 10
	xpto(&b)
	fmt.Println(b)
}