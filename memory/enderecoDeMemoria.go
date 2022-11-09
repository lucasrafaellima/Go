package main

import "fmt"

func main() {
	x := 10
	
	fmt.Println(&x)
	
	y := &x

	fmt.Println(y)
	fmt.Println(*y)
	*y = 20
	fmt.Println(x)
	var z *int = &x
	fmt.Println(z)
	fmt.Println(*z)
}