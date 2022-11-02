package main

import "fmt"

func main() {
	slice := make([]int, 5)
	for i := 0; i < len(slice); i++ {
		slice[i] += i+i
	}
	fmt.Println(slice)
}