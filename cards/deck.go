package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type deck,
// which is a slice of strigs
type deck []string // kinda extending []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) typeToString() string {
	return strings.Join([]string(d), "\n")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.typeToString()), 0644)
}

func readDeckFromFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(content), "\n")
}

func (d deck) shuffle() {
	lengthOfDeck := len(d) - 1
	//rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// find a random rumber between 0 and 51
		// interchange with current card in the random position
		newPosition := r.Intn(lengthOfDeck)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
