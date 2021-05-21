package main

import (
	"os"
	"reflect"
	"testing"
)

var fullOrderedDeck = deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"}

var	expectedDeckLength = len(fullOrderedDeck)
var	expectedFirstCard = fullOrderedDeck[0]
var	expectedLastCard = fullOrderedDeck[expectedDeckLength - 1]

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of %v, but got %v", expectedFirstCard, d[0])
	}

	if d[len(d) - 1] != expectedLastCard {
		t.Errorf("Expected last card of %v, but got %v", expectedLastCard, d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected %v cards in deck, got %v", expectedDeckLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T){
	expectedHand := fullOrderedDeck[:3]
	expectedStack := fullOrderedDeck[3:]

	d := newDeck()
	resultHand, resultStack := deal(d, 3)

	if !reflect.DeepEqual(expectedHand, resultHand) {
		t.Errorf("Expected hand %v, got %v", expectedHand, resultHand)
	}

	if !reflect.DeepEqual(expectedStack, resultStack) {
		t.Errorf("Expected stack %v, got %v", expectedStack, resultStack)
	}
}