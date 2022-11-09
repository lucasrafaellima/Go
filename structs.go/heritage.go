package main

import "fmt"

type plataform struct {
	xbox bool
	playStation bool
	nintendoSwift bool
}

type game struct {
	name           string
	year           string
	classification string
	plataform
}

func (g game) resp() string {
	return "name of game: " + g.name + "\ndate: " + g.year + "\nclassification: " + g.classification
}

func main() {
	gaming := game{"Call Of Duty Warzone", "2020", "18", plataform{xbox: true, playStation: false, nintendoSwift: false}}
	fmt.Println(gaming.resp())
}