package main

import "fmt"

func main() {
	cards := NewDeck()
	// firstHand, remaining := Deal(cards, 5)
	// fmt.Println("First Hand: ")
	// firstHand.Print()
	// fmt.Println("Remaining deck: ")
	// remaining.Print()
	fmt.Println(cards.toString())
	cards.saveToFile("myDeck")
}
