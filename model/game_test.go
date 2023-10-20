package model

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestNewGameEasyMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game := NewGame(playerName, Easy)
	assert.Equal(playerName, game.playerName)
	assert.Equal(easyDifficultyCardsTotal, len(game.cards))
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard)
}

func TestNewGameMediumMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game := NewGame(playerName, Medium)
	assert.Equal(playerName, game.playerName)
	assert.Equal(mediumDifficultyCardsTotal, len(game.cards))
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard)
}

func TestNewGameHardMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game := NewGame(playerName, Hard)
	assert.Equal(playerName, game.playerName)
	assert.Equal(HardDifficultyCardsTotal, len(game.cards))
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard)
}
