package main

import "fmt"

func main() {
	var x [10] int
	i := 0
	for i < len(x) {
		x[i] += i
		i += 1
	}

	fmt.Println(x)
}