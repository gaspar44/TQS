package model

import (
	"gaspar44/TQS/model/custom_errors"
	"math/rand"
	"time"
)

const (
	EasyDifficultyCardsTotal     = 6
	MediumDifficultyCardsTotal   = 10
	HardDifficultyCardsTotal     = 16
	easyDifficultyPenalization   = 1
	mediumDifficultyPenalization = 3
	hardDifficultyPenalization   = 5
)

var (
	initializationCard selectedCard
)

// Function to create/inizialate a Game
func NewGame(playerName string, gameDifficulty Difficulty) (*Game, error) {
	// TODO: start the points and initialize the remaining stuff
	ranking, err := GetRankingInstance()

	if err != nil {
		return nil, err
	}
	defaultCard := NewCard(-1)
	initializationCard = selectedCard{
		Card:     &defaultCard,
		Position: -1,
	}

	timerChannel := make(chan bool)
	createdGame := &Game{
		playerName:    playerName,
		difficulty:    gameDifficulty,
		Ranking:       ranking,
		points:        0,
		selectedCard:  initializationCard,
		endingChannel: timerChannel,
	}

	err = createdGame.createCards(gameDifficulty)
	return createdGame, err
}

type selectedCard struct {
	Card     *Card
	Position int
}

type Game struct {
	cards         []Card
	initialized   bool
	difficulty    Difficulty
	Ranking       *Ranking
	playerName    string
	points        int
	selectedCard  selectedCard
	endingChannel chan bool
}

func (g *Game) ChooseCardOnBoard(cardToSelect int) (bool, error) {
	if cardToSelect < 0 || cardToSelect > len(g.cards)-1 {
		return false, custom_errors.NewInvalidPositionError(cardToSelect)
	}

	previousSelectedCard := g.selectedCard
	card := &g.cards[cardToSelect]

	if card.IsDisable {
		return false, nil
	}

	card.Click()
	newSelectedCard := selectedCard{
		Card:     card,
		Position: cardToSelect,
	}

	if previousSelectedCard.Position != -1 && previousSelectedCard.Card.GetValue() != -1 && newSelectedCard.Card.GetValue() != previousSelectedCard.Card.GetValue() {
		switch g.difficulty {
		case Easy:
			g.points += easyDifficultyPenalization
		case Medium:
			g.points += mediumDifficultyPenalization
		case Hard:
			g.points += hardDifficultyPenalization
		}

		g.selectedCard.Card.Click()
		g.selectedCard = newSelectedCard
		return false, nil
	} else if newSelectedCard.Card.GetValue() == previousSelectedCard.Card.GetValue() && previousSelectedCard.Position == cardToSelect {
		// Same card selected. No penalization
		return true, nil
	} else if newSelectedCard.Card.GetValue() == previousSelectedCard.Card.GetValue() && newSelectedCard.Position != previousSelectedCard.Position {
		// If the both cards were correctly selected there is no penalization and the selected card is reset
		previousSelectedCard.Card.disable()
		card.disable()
		g.selectedCard = initializationCard
		return true, nil
	}
	g.points += 1
	g.selectedCard = newSelectedCard
	return true, nil
}

func (g *Game) Stop() {
	player := Player{
		Name:   g.playerName,
		Points: g.points,
	}

	g.Ranking.Update(player)
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
		for i := 0; i < EasyDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Medium:
		for i := 0; i < MediumDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Hard:
		for i := 0; i < HardDifficultyCardsTotal/2; i++ {
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
