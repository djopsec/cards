package main

import (
	"time"
)

func main() {

	cards := newDeck()
	cards.shuffle()
	
	hand, remainingCards := deal(cards, 3)

	hand.print()

	var t string = time.Now().Format(time.RFC3339)
	
	handFile := ("spread " + t)
	remainingFile := ("remaining cards " + t)

	hand.saveToFile(handFile)
	remainingCards.saveToFile(remainingFile)

}
