package main

import "testing"

func TestCreateNewDeck(t *testing.T) {

	d := createNewDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but get length of %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Error on First Element , Expected Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Error on Last Element , Expected King of Clubs but got %v", d[len(d)-1])
	}

}
