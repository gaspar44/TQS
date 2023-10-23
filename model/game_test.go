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
	assert.Equal(-1, game.selectedCard)

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
	assert.Equal(-1, game.selectedCard)

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
	assert.Equal(-1, game.selectedCard)

	expectedElements := make(map[int]int)

	for i := 0; i < hardDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
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
