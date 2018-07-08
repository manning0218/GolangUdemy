package main

import (
    "testing"
    "os"
)


func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) != 52 {
        t.Errorf("Expected deck length of 52, but got: %d", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("Expected 'Ace of Spades', but got: %s", d[0])
    }

    if d[len(d) - 1] != "King of Diamonds" {
        t.Errorf("Expected 'King of Diamonds', but got: %s", d[len(d) - 1])
    }
}

func TestSavingDeckToFile(t *testing.T) {
    d := newDeck()

    os.Remove("_decktesting.txt")
   
    d.saveToFile("_decktesting.txt")

    loadedDeck := newDeckFromFile("_decktesting.txt")

    if len(loadedDeck) != 52 {
        t.Errorf("Expected 52 cards, but got: %d", len(d))
    } else {
        os.Remove("_decktesting.txt")
    }
}
