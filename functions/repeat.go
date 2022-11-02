package main

func variadicFunc(x ...int) int {
	var anwser int
	for _, v := range x {
		anwser += v
	}
	return anwser

}

func main() {
	a := 6
	b := 10
	c := 21
	variadicFunc(a,b,c)
	
}