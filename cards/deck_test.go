package main

import (
	"os"
	"testing"
)

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

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {

	err := os.Remove("_testFile")
	if err != nil {

		newDeck := createNewDeck()
		newDeck.saveToFile("_testFile")

		newReadDeck := newDeckFromFile("_tetFile")

		if len(newReadDeck) != 52 {
			t.Errorf("Expected length of Deck is 52 but got %v", len(newReadDeck))
		}

		os.Remove("_testFile")
	}
}
