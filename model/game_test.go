package model

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	// Initialization of assert and playername
	print("Initialization of assert and player name...")
	assert := assert2.New(t)
	playerName := "test1"

	// Checking initialization of game
	print("Checking game...")
	game := NewGame(playerName, Easy)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard)

	print("End of card test!")
	// Missing: Display tests passed, cards assignation (?)
}
