package main 

import "fmt"

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	a := funcInsideFunc()
	fmt.Println(a)
}