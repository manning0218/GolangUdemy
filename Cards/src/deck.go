package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "math/rand"
    "time"
)

// Create a new type of 'deck' 
// This consists of a slice of strings
type deck []string

// This will create and return a list of cards
func newDeck() deck {
    cards := deck{}

    cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

    // Replace variable with underscore if not needed
    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }

    return cards
}

func deal(d deck, handSize int) (deck, deck){
    return d[:handSize], d[handSize + 1:]
}

// d deck is a receiver, now we can use access print() from specified type
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i + 1, card)
    }
}

// Helper function to convert a deck to a string
func (d deck) toString() string {

    return strings.Join([]string(d), ",")
}

// Saves file to disk
func (d deck) saveToFile(fileName string) error {
    return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Reads a file from the disk
func newDeckFromFile(fileName string) deck {
    // We receive a byteSlice and an error
    // error will have nil if no error occurred
    bs, err := ioutil.ReadFile(fileName)
    cards := deck{}

    if err != nil {
        fmt.Println("Error:", err)
        // This can exit the program entirely. Should occur if a fatal error has occurred
        // os.Exit(-1)
    } else {
        cards = deck(strings.Split(string(bs), ","))
    }

    return cards
}

func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    // Shuffle the deck 100 times
    for i := 0; i < 100; i++ {
        for j := range d {
            newPosition := r.Intn(len(d) - 1)

            d[j], d[newPosition] = d[newPosition], d[j]
        }
    }
}
