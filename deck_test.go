package main

import (
	"testing"
)

func Test_newDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards. Got %v", len(cards))
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", cards[len(cards)-1])
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Acec of Spades, but got %v", cards[0])
	}
}
