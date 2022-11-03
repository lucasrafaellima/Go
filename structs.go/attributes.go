package main

import "fmt"

type car struct {
	name string
	year int
	color string
}

func (c car) info() string {
	return c.name + " - " + c.color
}

func main() {
	car1 := car{"seda", 2000, "preto"}
	car2 := car{"siena", 2020, "azul"}
	fmt.Println(car1.color)
	fmt.Println(car2.name)
	fmt.Println(car1.info())
	fmt.Println(car2.info())
}