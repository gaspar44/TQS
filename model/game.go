package model

import "gaspar44/TQS/model/custom_errors"

const (
	easyDifficultyCardsTotal   = 6
	mediumDifficultyCardsTotal = 10
	hardDifficultyCardsTotal   = 16
)

type Game struct {
	cards        []Card
	initialized  bool
	difficulty   Difficulty
	Ranking      *Ranking
	playerName   string
	timer        int
	selectedCard int
}

func (g *Game) ChooseCardOnBoard() {
	// TODO
}

func (g *Game) GetCards() []Card {
	return g.cards
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
	return nil
}

// Function to create/inizialate a Game
func NewGame(playerName string, gameDifficulty Difficulty) (*Game, error) {
	// TODO: start the timer and initialize the remaining stuff
	createdGame := &Game{
		playerName:   playerName,
		difficulty:   gameDifficulty,
		Ranking:      nil,
		timer:        0,
		selectedCard: -1,
	}

	err := createdGame.createCards(gameDifficulty)

	return createdGame, err
}
