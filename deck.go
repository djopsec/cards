package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)
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
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, spreadSize int) (deck,deck) {
	return d[:spreadSize], d[spreadSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}