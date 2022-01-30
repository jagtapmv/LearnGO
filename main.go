package main

import "fmt"

func main() {
	cards := newCard()

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() []string {
	return []string{"Ace of Hearts", "Five of Diamonds", "King of Spades"}
}
