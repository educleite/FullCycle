package main

import "fmt"

type Item struct {
	ID          int
	code        string
	name        string
	description string
	unitValue   float64
}

func (item *Item) getItemName() string {

	return item.name

}

func (item *Item) setItemName(name string) {

	item.name = name

}

func main() {

	item := Item{
		ID:          1,
		code:        "IT001",
		name:        "Notebook DELL - standard",
		description: "Notebook DELL - INSPIRON 3000",
		unitValue:   3000,
	}

	// marcar a variavel ponteiro do struct Item
	// Ambas functios set e get sao metodos do struct
	item.setItemName("Notebook DELL - 17 polegadas - CORE i7")

	fmt.Print(item.getItemName())

}
