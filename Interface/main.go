package main

import (
    "fmt"
)

// Creates interface
type bot interface {
    getGreeting() string
}

// Used to be able to create functions tied
// to each struct and not worry about fields
type englishBot struct{}
type spanishBot struct{}

func main() {
    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

/*** ERROR: Cannot have functions with identical names
func printGreeting(eb englishBot) {
    fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
    fmt.Println(sb.getGreeting())
}
***/

func (eb englishBot) getGreeting() string {
    // Custom logic for english bot
    return "Hello!"
}

func (sb spanishBot) getGreeting() string {
    // Custom logic for spanish bot
    return "Hola!"
}
