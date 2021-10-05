package main

func main() {
	cards := newDeck()
	cards.saveToFile("deck.txt")
}

func newCard() string {
	return "Five of Diamonds"
}
