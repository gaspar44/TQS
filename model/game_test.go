package model

import (
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
	assert.Equal(0, game.timer)
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
	assert.Equal(0, game.timer)
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
	assert.Equal(0, game.timer)
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
	assert.Equal(0, game.timer)
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
	assert.Equal(0, game.timer)
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
	assert.Equal(0, game.timer)
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

func TestGameEasyModeChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(easyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard.Card.value)
	assert.False(game.selectedCard.Card.isVisible)
	assert.False(game.selectedCard.Card.isDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	err = game.ChooseCardOnBoard(0)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
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
