package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// `cards` can call `print` because it is of type `deck`.
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// Create a new deck from file.
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// Shuffle cards.
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
