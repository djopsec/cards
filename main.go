package main

// import "fmt"

func main() {
	// cards := newDeck()

	// spread, remainingCards := deal(cards, 5)

	// spread.print()
	// remainingCards.print()
	cards := newDeck()
	cards.saveToFile("my_cards")
}
