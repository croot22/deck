package cardDeck

import (
	"math/rand"
	"time"
)

type Card struct{
	number int
	suit string
}

func main(){
	deck := createDeck()
	shuffleDeck(deck)
}

func createDeck(numJokers int) []Card{
	var deck []Card
	for i := 1; i < 5; i++{
		for j := 1; j < 14; j++{
			var currCard = Card{}
			if i == 1{
				currCard.suit = "s"
			} else if i == 2{
				currCard.suit = "d"
			} else if i == 3{
				currCard.suit = "c"
			} else if i == 4{
				currCard.suit = "h"
			}
			currCard.number = j
			deck = append(deck, currCard)
		}
	}
	if numJokers > 0{
		for i := 0; i < numJokers; i++{
			joker := Card{
				number: 0,
				suit: "j",
			}
			deck = append(deck, joker)
		}
	}
	return deck
}

func shuffleDeck(deck []Card){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {deck[i], deck[j] = deck[j], deck[i]})
}

func sortDeck(deck []Card, hands, numCardsPerHand int, hasDraw bool){

}