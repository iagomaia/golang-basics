package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("deck.txt")
	readCards := newDeckFromFile("deck.txt")
	readCards.print()
}
