package model

type Game struct {
	cards        *[]Card
	Ranking      *Ranking
	playerName   string
	timer        int
	selectedCard int
}

func (g *Game) ChooseCard() {
	// TODO
}

func NewGame(playerName string, gameDifficulty Difficulty) *Game {
	// TODO: start the timer and initialize the remaining stuff

	return &Game{
		playerName:   playerName,
		cards:        nil,
		Ranking:      nil,
		timer:        0,
		selectedCard: -1,
	}
}

func createCards(difficulty Difficulty) *[]Card {
	return nil
}
