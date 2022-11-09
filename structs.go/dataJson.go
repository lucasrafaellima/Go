package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

type Dates struct {
	Year1 string
	Year2 string
	Year3 string
}

func main() {
	car2 := Car{"bmw", 2021, "black"}
	result, _ := json.Marshal(car2)
	fmt.Println(string(result))
}