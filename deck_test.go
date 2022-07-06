package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 78 {
		t.Errorf("Expected deck length of 78, but got %v", len(d))
	}

	if d[0] != "00-The Fool" {
		t.Errorf("Expected first card of 00-The Fool, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Pentacles" {
		t.Errorf("Expected last card of King of Pentacles, but got %v", d[len(d)-1])
	}

	
}