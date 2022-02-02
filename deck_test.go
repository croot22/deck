package cardDeck

import (
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
	sortDeck(deck, 4, 52/4 - 4, true)

}