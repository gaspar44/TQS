package model

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game := NewGame(playerName, Easy)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.timer)
	assert.Equal(-1, game.selectedCard)
}
