package model

//////////////////////////////////////////////////////////////////////////
type Game struct {
	cards        *[]Card
	Ranking      *Ranking
	playerName   string
	timer        int
	selectedCard int
}

//////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////
// Function to choose a board card
func (g *Game) ChooseCard() {
	// TODO
}

//////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////
// Function to create/inizialate a Game
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

//////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////
// Function to create cards
func createCards(game *Game, difficulty Difficulty) {

	switch difficulty {
	case Easy:
		for i := 1; i <= 5; i++ {
			newCard := NewCard(i)
			game.cards = append(game.cards, newCard)
		}
	case Medium:
		for i := 1; i <= 10; i++ {
			newCard := NewCard(i)
			game.cards = append(game.cards, newCard)
		}
	case Hard:
		for i := 1; i <= 15; i++ {
			newCard := NewCard(i)
			game.cards = append(game.cards, newCard)
		}
	}

}

//////////////////////////////////////////////////////////////////////////
