package main

import (
	"testing"
)

func Test_shouldTakeHigher(t *testing.T) {

	existingDeck := []*Card{newCard(2, "Spades"), newCard(4, "Spades"), newCard(4, "Diamonds")}
	if !shouldTakeHigher(existingDeck, 3) {
		t.Errorf("exected to take higher")
	}

}

func Test_newDeck(t *testing.T) {

	// test ace high
	aceHighDeck := *newDeck(true)

	sizeOfNewDeck := len(aceHighDeck)

	if sizeOfNewDeck != 52 {
		t.Errorf("expected size of new deck to be 52 but was %d", sizeOfNewDeck)
	}

	thirdCard := aceHighDeck[2]
	if (thirdCard.Number) != 4 {
		t.Errorf("expected the third card to have value 4 but was %d (%s)", thirdCard.Number, thirdCard)
	}

	twentiethCard := aceHighDeck[19]
	if (twentiethCard.Number) != 8 {
		t.Errorf("expected the twentieth card to have value 8 but was %d (%s)", twentiethCard.Number, twentiethCard)
	}

	// test ace low
	aceLowDeck := *newDeck(false)

	firstAceLowCard := aceLowDeck[0]
	if firstAceLowCard.Number != 1 {
		t.Errorf("expected an ace to be the first card in an ace low set but the first card was %d (%s)", firstAceLowCard.Number, firstAceLowCard)
	}

}

func Test_WinsIteration(t *testing.T) {
	fourOfSpades := newCard(4, "Spades")
	fiveOfDiamonds := newCard(5, "Diamonds")

	// going higher
	if WinsIteration(true, fourOfSpades, fiveOfDiamonds) == false {
		t.Error("Expected to win with prev 4, next 5 and user going higher")
	}

	if WinsIteration(true, fiveOfDiamonds, fourOfSpades) == true {
		t.Error("Expected to lose with prev 5, next 4 and user going higher")
	}

	if WinsIteration(true, fiveOfDiamonds, fiveOfDiamonds) == false {
		t.Error("Expected to win with prev 5, next 5 and user going higher")
	}

	// going lower
	if WinsIteration(false, fourOfSpades, fiveOfDiamonds) == true {
		t.Error("Expected to win with prev 4, next 5 and user going lower")
	}

	if WinsIteration(false, fiveOfDiamonds, fourOfSpades) == false {
		t.Error("Expected to win with prev 5, next 4 and user going lower")
	}

	if WinsIteration(false, fiveOfDiamonds, fiveOfDiamonds) == false {
		t.Error("Expected to lose with prev 5, next 5 and user going lower")
	}

}
