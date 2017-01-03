package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	totalIterations := flag.Int("iterations", 1000000, "number of iterations")
	aceHigh := flag.Bool("ace-high", false, "ace high?")
	verbose := flag.Bool("v", false, "verbose? (show logging)")
	flag.Parse()

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	winningCount := 0

	for iteration := 0; iteration < *totalIterations; iteration++ {
		deck := newDeck(*aceHigh)
		deck.Shuffle()
		log.Printf("Shuffled the deck. %v\n", deck)
		wins := deck.WinsRun()
		if wins {
			winningCount++
		}
		log.Printf("wins? %t\n", wins)
	}

	fmt.Printf("Won %d out of %d times. %f%%\n", winningCount, *totalIterations,
		float64(winningCount)*100/float64(*totalIterations))

}

// wins the whole run (ie, all 12 cards guessed correctly)
func (deckPtr *Deck) WinsRun() bool {

	deck := *deckPtr

	for numberOfCardsRevealed := 1; numberOfCardsRevealed < 12; numberOfCardsRevealed++ {

		previousCard := deck[numberOfCardsRevealed-1]
		remainingDeck := deck[numberOfCardsRevealed:]

		userShouldChooseHigher := shouldTakeHigher(remainingDeck, previousCard.Number)

		nextCard := deck[numberOfCardsRevealed]

		log.Printf("%d cards revealed. previous card: %s, picking to go higher? %t\n", numberOfCardsRevealed, previousCard, userShouldChooseHigher)

		if !WinsIteration(userShouldChooseHigher, previousCard, nextCard) {
			return false
		}

	}
	return true

}

// WinsIteration sees whether the player wins one single card turnover
// if the number is in the same direction they guessed in, or the same as the previous number, the player wins.
func WinsIteration(userShouldChooseHigher bool, previousCard *Card, nextCard *Card) bool {

	if userShouldChooseHigher && nextCard.Number < previousCard.Number {
		return false
	}
	if !userShouldChooseHigher && nextCard.Number > previousCard.Number {
		return false
	}
	return true
}

func shouldTakeHigher(remainingDeck []*Card, previousCardValue int) bool {
	amountOfCardsHigher := 0
	amountOfCardsLower := 0
	for _, card := range remainingDeck {
		if card.Number > previousCardValue {
			amountOfCardsHigher++
		} else if card.Number < previousCardValue {
			amountOfCardsLower++
		}
	}

	return amountOfCardsHigher > amountOfCardsLower

}

type Card struct {
	Number int
	Suit   string
}

type Deck []*Card

func (card *Card) String() string {
	var numberAsString string
	switch card.Number {
	case 1, 14:
		numberAsString = "Ace"
	case 13:
		numberAsString = "King"
	case 12:
		numberAsString = "Queen"
	case 11:
		numberAsString = "Jack"
	default:
		numberAsString = strconv.Itoa(card.Number)
	}

	return fmt.Sprintf("%s of %s", numberAsString, card.Suit)
}

func newCard(number int, suit string) *Card {
	return &Card{
		Number: number,
		Suit:   suit,
	}
}

// sorted deck
func newDeck(aceHigh bool) *Deck {
	var deck Deck
	suits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}

	startNumber := 1 // ace low
	endNumber := 13  // king
	if aceHigh {
		startNumber = 2
		endNumber = 14 // ace high
	}

	for _, suit := range suits {
		for number := startNumber; number <= endNumber; number++ {
			deck = append(deck, newCard(number, suit))
		}
	}
	return &deck
}

func (deckPtr *Deck) Shuffle() {
	deck := *deckPtr
	rand.Seed(time.Now().UnixNano())

	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func (deckPtr *Deck) String() string {
	str := ""
	for index, card := range *deckPtr {
		if index != 0 {
			str += ", "
		}
		str += fmt.Sprintf("[%s]", card)
	}
	return str
}
