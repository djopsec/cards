package main

func main() {
	cards := newDeck()

	spread, remainingCards := deal(cards, 5)

	spread.print()
	remainingCards.print()
}
