package model

const (
	easyDifficultyCardsTotal   = 6
	mediumDifficultyCardsTotal = 10
	HardDifficultyCardsTotal   = 16
)

type Game struct {
	cards        []Card
	difficulty   Difficulty
	Ranking      *Ranking
	playerName   string
	timer        int
	selectedCard int
}

func (g *Game) ChooseCardOnBoard() {
	// TODO
}

func createCards(difficulty Difficulty) []Card {
	cards := make([]Card, 0)

	switch difficulty {
	case Easy:
		for i := 0; i < easyDifficultyCardsTotal; i++ {
			newCard := NewCard(i)
			cards = append(cards, newCard)
		}
	case Medium:
		for i := 0; i < mediumDifficultyCardsTotal; i++ {
			newCard := NewCard(i)
			cards = append(cards, newCard)
		}
	case Hard:
		for i := 0; i < HardDifficultyCardsTotal; i++ {
			newCard := NewCard(i)
			cards = append(cards, newCard)
		}
	}
	return cards
}

// Function to create/inizialate a Game
func NewGame(playerName string, gameDifficulty Difficulty) *Game {
	// TODO: start the timer and initialize the remaining stuff
	cards := createCards(gameDifficulty)
	return &Game{
		playerName:   playerName,
		difficulty:   gameDifficulty,
		cards:        cards,
		Ranking:      nil,
		timer:        0,
		selectedCard: -1,
	}
}
