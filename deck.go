package main

import "fmt"

type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
