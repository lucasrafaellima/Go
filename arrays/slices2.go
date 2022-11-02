package main

import "fmt"

func main() {
	slice2 := make([]int, 1)
	slice2[0] = 11
	fmt.Println(slice2)
	slice2 = append(slice2, 4,6,7,3,1)
	fmt.Println(slice2)
}