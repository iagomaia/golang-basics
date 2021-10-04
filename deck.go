package main

import "fmt"

type deck [] string;

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card);
	}
}

func newDeck() deck {
	
}