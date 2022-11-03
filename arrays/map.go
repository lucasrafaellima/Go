package main

import "fmt"

func main() {
	m := make(map[string] int)
	m["a"] = 19
	m["b"] = 20

	delete(m, "a")
	_, exists := m["a"]
	_, exists2 := m["b"]
	fmt.Println(exists)
	fmt.Println(exists2)
}