package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) getSingleString() string {
	return strings.Join(d, "\n")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.getSingleString()), 0666)
}

func readFromFile(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return b
}

func newDeckFromFile(fileName string) deck {
	bs := readFromFile(fileName)
	deckString := string(bs)
	stringSlice := strings.Split(deckString, "\n")
	cards := []string{}
	cards = append(cards, stringSlice...)
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)

	for i := range d {
		newPosition := rng.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
