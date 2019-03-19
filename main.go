package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card struct defines the card itself
type Card struct {
	Suit      string
	SuitValue int
	Rank      int
	RankValue string
}

// Player struct contains a card
type Player struct {
	Cards      []Card
	TotalScore int
	IsDealer   bool
}

// Game contains a arbitary number of players
type Game struct {
	Players []Player
}

// Deck struct contains a list of cards
type Deck []Card

var suitList = map[string]int{"Spade": 1, "Heart": 2, "Club": 3, "Diamond": 4}
var rankList = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Jack",
	11: "Queen",
	12: "King",
	13: "Ace",
}

// Initialize creates the necessary player's card
func Initialize(game *Game, numOfPlayers int, deck *Deck) (Game, Deck) {
	// create a dealer first
	player := Player{}
	player.IsDealer = true
	game.Players = append(game.Players, player)

	// create the rest of the players
	for playerNum := 0; playerNum < numOfPlayers; playerNum++ {
		player := Player{}
		game.Players = append(game.Players, player)
	}

	for dealCount := 0; dealCount < 2; dealCount++ {
		for index, player := range game.Players {
			player, *deck = Deal(&player, deck)
			game.Players[index] = player
		}

	}
	return *game, *deck
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	deck := Deck{}

	for suit, suitValue := range suitList {
		for rank, rankValue := range rankList {
			card := Card{
				Suit:      suit,
				SuitValue: suitValue,
				Rank:      rank,
				RankValue: rankValue,
			}
			deck = append(deck, card)
		}
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	game := Game{}
	game, deck = Initialize(&game, 1, &deck)

	for index, player := range game.Players {
		player.TotalScore = CheckValue(player)
		game.Players[index] = player
	}

	for index, player := range game.Players {
		if index > 0 {
			fmt.Println(fmt.Sprintf("Player %v has following cards:", index+1))
			for _, card := range player.Cards {
				fmt.Println(fmt.Sprintf("%v of %v", card.Suit, card.RankValue))
			}
		}
	}

	fmt.Println("Dealer has the following cards:")
	dealer := game.Players[0]
	for _, card := range dealer.Cards {
		fmt.Println(fmt.Sprintf("%v of %v", card.Suit, card.RankValue))
	}

	for index, player := range game.Players {
		if index > 0 {
			if player.TotalScore == 0 {
				fmt.Println("Player one blackjack and wins automatically")
			} else if dealer.TotalScore == 0 {
				fmt.Println("Player two blackjack and wins automatically")
			} else if dealer.TotalScore > player.TotalScore {
				fmt.Println("Dealer wins")
			} else if player.TotalScore > dealer.TotalScore {
				fmt.Println(fmt.Sprintf("Player %v wins", index+1))
			} else {
				fmt.Println("Draw")
			}
		}
	}
}

// Deal add a card to a given player
func Deal(player *Player, deck *Deck) (Player, Deck) {
	player.Cards = append(player.Cards, (*deck)[0])
	return *player, (*deck)[1:]
}

func compare(cardOne Card, cardTwo Card) int {
	if cardOne.RankValue > cardTwo.RankValue {
		return 1
	} else if cardOne.RankValue < cardTwo.RankValue {
		return -1
	} else {
		if cardOne.SuitValue > cardTwo.SuitValue {
			return 1
		} else if cardOne.SuitValue < cardTwo.SuitValue {
			return -1
		}
	}

	return 0
}

// CheckValue computes the points for a given player's card
func CheckValue(player Player) int {
	cards := player.Cards
	totalScore := 0
	for index := range cards {
		totalScore += cards[index].Rank
	}

	// check blackjack
	result := CheckBlackJack(cards[0], cards[1])

	if result {
		return 0
	}

	return totalScore
}

// CheckBurst checks whether the player's cards has burst 21 points
func CheckBurst(player Player) bool {
	totalScore := 0
	actualPointsMap := map[int]int{
		11: 10,
		12: 10,
		13: 10,
	}

	for _, card := range player.Cards {
		if card.Rank > 10 {
			totalScore += actualPointsMap[card.Rank]
		} else {
			totalScore += card.Rank
		}
	}

	return totalScore > 21
}

// CheckBlackJack check if the combination is a blackjack
func CheckBlackJack(cardOne Card, cardTwo Card) bool {
	if cardOne.Rank == 14 && cardTwo.Rank >= 10 && cardTwo.Rank <= 12 {
		return true
	} else if cardTwo.Rank == 14 && cardOne.Rank >= 10 && cardOne.Rank <= 12 {
		return true
	}
	return false
}
