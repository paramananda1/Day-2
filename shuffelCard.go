package main

import (
	"fmt"
	"math/rand"
)

// Card holds the card suits and types in the deck
type Card struct {
	Type string
	Suit string
}

// Deck holds the cards in the deck to be shuffled
type Deck []Card

// New creates a deck of cards to be used
func NewDeck() (deck Deck) {

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	// appending to the deck
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

// Shuffle the deck
func shuffle_Main(d Deck , breaknumber int) Deck {
	// swap decks from breaknumber
var tempdeck Deck

	for i:=breaknumber;i< len(d);i++{
		tempdeck = append(tempdeck, d[i])
	}
	for i:=0;i< breaknumber;i++{
		tempdeck = append(tempdeck, d[i])
	}
	//fmt.Println(tempdeck)
	return tempdeck
}


func Shuffle() {
	deck := NewDeck()
	fmt.Println("Before Shuffel deck:")
	fmt.Println(deck)
	r := rand.Intn(52)
	fmt.Println("random number",r)

	d1 := shuffle_Main(deck,r)
	fmt.Println("After Shuffel deck:")
	fmt.Println(d1)

	fmt.Println(len(deck), len(d1))

}