package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //converted deck d to []string and used Join to join each element of []string using seperator ","
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func deckFromDrive(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	sDeck := strings.Split(string(data), ",")
	return deck(sDeck)
}

func (d deck) shuffle() deck {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	source := rand.NewSource(int64(binary.LittleEndian.Uint64(b[:])))
	r := rand.New(source)

	for i := range d {
		randCardPosition := r.Intn(len(d))
		d[i], d[randCardPosition] = d[randCardPosition], d[i]
	}
	return d
}
