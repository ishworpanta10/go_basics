package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is a slice of strings
type deck []string

func createNewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	// sliceOfString := []string(d)

	// singleString := strings.Join(sliceOfString, ",")

	// return singleString

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(fileName string) error {

	stringData := d.toString()

	return os.WriteFile(fileName, []byte(stringData), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {

		// Option #1 - log error and return a call to newDeck
		// Option #2 - log error and exit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert byte slice to deck
	stringData := string(bs)
	stringList := strings.Split(stringData, ",")
	return deck(stringList)
}

func (d deck) shuffle() {
	// own random seed from time
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// receiver function for deck type that prints all the cards in the deck
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)

	}
}
