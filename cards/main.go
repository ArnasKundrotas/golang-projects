package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard() // := new variable assignment
	//card = "Five of Diamonds" = for existing variables no :
	fmt.Println(card) // import "fmt" on save
}

func newCard() string {
	return "Five of Diamonds"
}
