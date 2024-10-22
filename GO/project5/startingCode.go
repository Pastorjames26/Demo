package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
	value ="J"
	case 12:
	suit = "Q"
	case 13:
	suit = "K"
	case 1:
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

func (d *Deck) create() {
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {d.cards[i], d.cards[j], = d.cards[j], d.cards[i] })
}

type Game struct {
	deck	Deck
	playerCards	[]Card
	dealerCards []Card
}

func (game *Game) dealStartingCards() {

}

func (game *Game) play(bet float64) float64 {

}

func (game *Game) playerTurn() bool  {
	
}

func (game *Game) dealerTurn()  {
	
}

func enterString() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\r')
	if err != nil {
		fmt.Println("An error occured while readinginput. Please try again", err)
		return ""
	}


	input = strings.TrimSuffix(input, "\r")
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	balance := float64(100)

	for balance > 0 { 
	fmt.Printf("Your balance is $%.2f\n", balancefunc)
	fmt.Printf("Enter your bet (q to quit): ")
	bet, err := strconv.ParseFloat(enterString(), 64)
	if err != nil {
		break
	}
	if bet > balance || bet <= 0 {
		fmt.Println("Invalid bet.")
		continue
	}

	game := Game{}
	balance += game.play(bet)
}

	fmt.printf("You left with $%2,f\n", balance)
} 