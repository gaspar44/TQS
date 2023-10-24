package model

import (
	"gaspar44/TQS/model/custom_errors"
	"math/rand"
	"time"
)

const (
	easyDifficultyCardsTotal     = 6
	mediumDifficultyCardsTotal   = 10
	hardDifficultyCardsTotal     = 16
	easyDifficultyPenalization   = 1
	mediumDifficultyPenalization = 3
	hardDifficultyPenalization   = 5
)

// Function to create/inizialate a Game
func NewGame(playerName string, gameDifficulty Difficulty) (*Game, error) {
	// TODO: start the timer and initialize the remaining stuff
	initializationCard := NewCard(-1)
	createdGame := &Game{
		playerName: playerName,
		difficulty: gameDifficulty,
		Ranking:    nil,
		timer:      0,
		selectedCard: selectedCard{
			Card:     initializationCard,
			Position: -1,
		},
	}

	err := createdGame.createCards(gameDifficulty)
	return createdGame, err
}

type selectedCard struct {
	Card     Card
	Position int
}

type Game struct {
	cards        []Card
	initialized  bool
	difficulty   Difficulty
	Ranking      *Ranking
	playerName   string
	timer        int
	selectedCard selectedCard
}

func (g *Game) ChooseCardOnBoard(cardToSelect int) error {
	if cardToSelect < 0 || cardToSelect > len(g.cards)-1 {
		return custom_errors.NewInvalidPositionError(cardToSelect)
	}

	newSelectedCard := g.cards[cardToSelect]
	newSelectedCard.Click()

	if newSelectedCard.GetValue() != g.selectedCard.Card.GetValue() {
		switch g.difficulty {
		case Easy:
			g.timer += easyDifficultyPenalization
		case Medium:
			g.timer += mediumDifficultyPenalization
		case Hard:
			g.timer += hardDifficultyPenalization
		}
		g.selectedCard.Card.Click()
	}

	g.selectedCard = selectedCard{
		Card:     newSelectedCard,
		Position: cardToSelect,
	}

	return nil
}

func (g *Game) updateTimer() {
	//
}

func (g *Game) GetCards() []Card {
	return g.cards
}

func (g *Game) shuffleCards() {
	if !g.initialized {
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(g.cards), func(i, j int) {
		g.cards[i], g.cards[j] = g.cards[j], g.cards[i]
	})
}

func (g *Game) createCards(difficulty Difficulty) error {
	cards := make([]Card, 0)

	if g.initialized {
		return &custom_errors.GameAlreadyInitializedError{}
	}

	switch difficulty {
	case Easy:
		for i := 0; i < easyDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Medium:
		for i := 0; i < mediumDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Hard:
		for i := 0; i < hardDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	}

	g.cards = cards
	g.initialized = true
	g.shuffleCards()
	return nil
}
