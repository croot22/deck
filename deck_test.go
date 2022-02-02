package cardDeck

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_createDeck(t *testing.T) {
	deck := createDeck(0)
	if len(deck) != 52{
		t.Errorf("Deck not created correctly")
	}

	deckWithJokers := createDeck(2)
	if len(deckWithJokers) != 54{
		t.Errorf("Deck not created correctly")
	}
}

func Test_shuffleDeck(t *testing.T) {
	shuffledDeck := createDeck(0)
	shuffleDeck(shuffledDeck)

	origDeck := createDeck(0)

	if reflect.DeepEqual(origDeck, shuffledDeck){
		t.Errorf("Deck was not shuffled")
	}
}

func Test_sort(t *testing.T) {
	deck := createDeck(0)
	numCardsPerHand := 11
	hand1 := make([]Card, numCardsPerHand)
	hand2 := make([]Card, numCardsPerHand)
	hands := [][]Card{hand1, hand2}
	sortDeck(deck, hands, numCardsPerHand, false)
	fmt.Printf("%v\n", hand1)
	fmt.Printf("%v\n", hand2)
	if len(hand1) != 11 || len(hand2) != 11 {
		t.Errorf("Deck not sorted correctly")
	}
}