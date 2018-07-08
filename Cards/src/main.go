package main

import "fmt"

func main() {
//    cards := newDeck()

    fmt.Println("New Deck:")
//    cards.saveToFile("DeckOfCards.txt")

    newDeck := newDeckFromFile("DeckOfCards.txt")
    newDeck.print()
    fmt.Println()

    newDeck.shuffle()
    newDeck.print()
}

