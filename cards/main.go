package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, _ := deal(cards, 5)
	hand.saveToFile("hand")
	fmt.Println(readDeckFromFile("hand"))
}
