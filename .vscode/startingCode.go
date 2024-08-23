package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	value int
	 suit int
}

func (card Card) getString() string {
	var suit string
	var value string

	switch card.suit {
	case 0:
		suit = ""
	case 1:
		suit = ""
	case 2:
		suit = ""
	case 3:
		suit = ""
	}

	switch card.value {
	case 11: 
		suit = "J"
	case 11: 
		suit = "Q"
	case 11: 
		suit = "K"
	case 11: 
		suit = "A"
	default:
		value = fmt.Sprintf("%d", card.value)
	}

	return value + suit
}

type Deck struct {
	cards []Card
}

func (d *Deck) deal(num uint) []Card {
}

func (d* Deck) create() {
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {d.cards[j] = d.cards[j], d.cards[i]})
}

type Game struct {
	deck 		Deck
	playerCards	[]Card
	dealerCards	[]Card
}

func (game *Game) dealStartingCards(){

}

func (game *Game) play(bet float64) float64 {

}

func (game *Game) playerTurn() bool {

}

func (game *Game) dealerTurn() {

}

func enterString() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\r') // if \r doesnt work use \n
	if err != nil{
		fmt.Println("An error occurred while reading input. Please try again", err)
		return ""

	}
	//remove the delimiter from the string
	input = strings.TrimSuffix(input, "\r")
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	balance := float64(100)

	for balance > 0 {
		fmt.Printf("your balance is: $%.2f\n", balance)
		fmt.Printf("Enter your bet (q to quit): ")
		bet, err := strconv.ParseFloat(enterString(), 64)
		if err != nil {
			break
		}
		if bet > balance || bet <= 0 {
			fmt.Println("invalid bet.")
			continue
		}

		game := Game{}
		balance += game.play(bet)
	}

	fmt.Printf("You left with: $%2.f\n" balance)
}