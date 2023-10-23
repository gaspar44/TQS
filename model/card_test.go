package model

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	// Initialization of assert and card
	assert := assert2.New(t)
	card := NewCard(1)

	// Checking card value (Needed here?)
	assert.Equal(1, card.value)

	// Checking card visibility
	assert.False(card.isVisible)

	// Checking card "Click"
	card.Click()
	// Checking card visibility
	assert.True(card.isVisible)
	// Missing: Display tests passed
}
