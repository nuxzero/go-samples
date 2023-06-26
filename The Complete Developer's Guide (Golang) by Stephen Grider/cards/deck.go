package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of `deck` which is a slice of strings.
type deck []string

// why it don't have receiver? Because we don't need to create a deck to create a new deck.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubes"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// `(d deck)` is a receiver. Any variable of type `deck` now gets access to the `print` method.
// Why `d`? By convention, the receiver is a one or two letter abbreviation of the type.
// In Go, we don't use `this` or `self` as in other languages.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returns two values of type `deck` from handSize value.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// `0666` means it can read and write file.
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // 1 means something bad happened.
	}

	// Parses slice of bytes to string.
	str := string(bs)

	// Splits string by comma to a slice of strings.
	s := strings.Split(str, ",")

	// Turns slice of strings to a deck.
	return deck(s)
}

func (d deck) shuffle() {
	// Creates a new source (seed) for random numbers.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// `rand.Intn` returns a random number between 0 and n-1.
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap values.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
