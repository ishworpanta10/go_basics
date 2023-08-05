package main

import "fmt"

func main() {
	cards := createNewDeck()
	// hand, remainingCards := deal(cards, 2)
	// cards := newDeckFromFile("test.txt")
	hand, remainingCards := deal(cards, 2)

	fmt.Println("=============== Hand Cards ==============")

	hand.print()

	fmt.Println("=============== Remaining Cards ==============")

	remainingCards.print()

	// fmt.Println(hand.toString())

	cards.shuffle()

	cards.print()

	// hand.saveToFile("test.txt")

	// newDeckFromFile("test.txt")

	// cards.print()

}
