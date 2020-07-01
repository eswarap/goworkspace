package main

func main() {

	anotherCards := newDeckFromFile("mycards")
	anotherCards.shuffle()
	anotherCards.print()
}
