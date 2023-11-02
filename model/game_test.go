package model

import (
	"gaspar44/TQS/model/custom_errors"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestNewGameEasyMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < easyDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

func TestNewGameMediumMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Medium)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(mediumDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < mediumDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

func TestNewGameHardMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(hardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.True(game.initialized)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < hardDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

func TestGameEasyModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Easy)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

func TestGameMediumModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Medium)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(mediumDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

func TestGameHardModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Hard)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(hardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

func TestGameEasyModeSingleChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
}

func TestGameEasyModeSameCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.Equal(-1, game.selectedCard.Position)

	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Equal(newSelectedCard, game.selectedCard)
}

func TestGameHardModeSingleChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(hardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(4)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
}

func TestGameEasyModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(4)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(easyDifficultyPenalization+1, game.points)
}

func TestGameMediumModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Medium)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(mediumDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Medium)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(5)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(mediumDifficultyPenalization+1, game.points)
}

func TestGameHardModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(hardDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Hard)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(5)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(hardDifficultyPenalization+1, game.points)
}

func TestGameHardModeChoseSameCardTwice(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(hardDifficultyCardsTotal, len(game.GetCards()))

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	sameCardChoice := game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.Equal(sameCardChoice, newSelectedCard)
	assert.Equal(1, game.points)
}

func TestGameCorrectCards(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(1)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	assert.Equal(initializationCard.Card.GetValue(), newSelectedCard.Card.GetValue())
	assert.Equal(initializationCard.Position, newSelectedCard.Position)
	assert.True(previousSelectedCard.Card.isDisable)
	assert.True(game.GetCards()[1].isDisable)
	assert.Equal(1, game.points)
}

func TestGameSelectDisabledCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.Nil(err)
	assert.True(correctCards)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(1)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	correctCards, err = game.ChooseCardOnBoard(1)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard

	assert.Equal(-1, newSelectedCard.Card.value)
	assert.False(newSelectedCard.Card.isVisible)
	assert.False(newSelectedCard.Card.isDisable)
	assert.Equal(-1, newSelectedCard.Position)
}

func TestGameStop(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)

	ranking, err := GetRankingInstance()
	assert.Nil(err)
	assert.NotNil(ranking)
	defer ranking.release()
	players := []Player{
		{
			Name:   "test1",
			Points: 0,
		},
		{
			Name:   "test2",
			Points: 10,
		},
	}

	ranking.SetPlayers(players)
	game.Stop()

	setPlayers, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.Equal(players, setPlayers)
}

func TestGameInvalidLowerCardSelection(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))

	invalidPosition := -1
	correctCards, err := game.ChooseCardOnBoard(invalidPosition)
	assert.False(correctCards)
	assert.NotNil(err)
	assert.Equal(err.Error(), custom_errors.InvalidCardPositionErrorMessage+"-1")
}

func createCards(difficulty Difficulty) []Card {
	cards := make([]Card, 0)

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

	return cards
}

func checkCardsOrder(beforeShuffleCards []Card, afterShuffleCards []Card) bool {
	for i := range beforeShuffleCards {
		if beforeShuffleCards[i].GetValue() != afterShuffleCards[i].GetValue() {
			return true
		}
	}
	return false
}

func countDuplicates(cards []Card) map[int]int {
	countedElements := make(map[int]int)

	for i := 0; i < len(cards); i++ {
		cardValue := cards[i].value
		_, existsKey := countedElements[cardValue]

		if !existsKey {
			countedElements[cardValue] = 1
		} else {
			countedElements[cardValue] += 1
		}
	}
	return countedElements
}
