package main

func main() {
	cards := newCard()
	cards.print()
}

func newCard() deck {
	return deck{"Ace of Hearts", "Five of Diamonds", "King of Spades"}
}
