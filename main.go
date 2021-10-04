package main


func main() {
	cards := deck {newCard(), "Ace of Diamonds"};
	cards = append(cards, "Six of Hearts");

	cards.print();
}

func newCard() string {
	return "Five of Diamonds";
}