package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cardTypes := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := []string{}

	for _, cardType := range cardTypes {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardType)
		}
	}
	return cards
}

func saveToFile() {
	
}