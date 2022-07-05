package main

import "fmt"
// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardMajor := []string{"00 - The Fool", "01 - The Magician", "02 - The High Priestess", "03 - The Empress", "04 - The Emperor", "05 - The Hierophant", "06 - The Lovers", "07 - The Chariot", "08 - Strength", "09 - The Hermit", "10 - Wheel of Fortune", "11 - Justice", "12 - The Hanged One", "13 - Death", "14 - Temperance", "15 - The Devil", "16 - The Tower", "17 - The Star", "18 = The Moon", "19 - The Sun", "20 - Judgement", "21 - The World"}
	cardSuits := []string{"Wands", "Cups", "Swords", "Pentacles"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Page", "Knight", "Queen", "King"}

	for _, major := range cardMajor{
		cards = append(cards, major)
	}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d{
		fmt.Println(card)
	}
}