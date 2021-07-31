package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("deck_of_cards.txt")
	cards := newDeckFromFile("deck_of_cards.txt")
	cards.print()
}
