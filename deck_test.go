package cardDeck

import (
	"reflect"
	"testing"
)

func Test_createDeck(t *testing.T) {
	deck := createDeck()
	if len(deck) != 52{
		t.Errorf("Deck not created correctly")
	}
}

func Test_shuffleDeck(t *testing.T) {
	shuffledDeck := createDeck()
	shuffleDeck(shuffledDeck)

	origDeck := createDeck()

	if reflect.DeepEqual(origDeck, shuffledDeck){
		t.Errorf("Deck was not shuffled")
	}
}

func Test_sort(t *testing.T) {
	deck := createDeck()
	sortDeck(deck, 4, 52/4 - 4, true)

	if
}