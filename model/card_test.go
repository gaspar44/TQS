package model

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	// Initialization of assert and card
	print("Initialization of assert and card...")
	assert := assert2.New(t)
	card := NewCard(1)
	print("Card value: ", card.value)
	print("Card visibility: ", card.isVisible)

	// Checking card value (Needed here?)
	print("Checking card value...")
	assert.Equal(1, card.value)

	// Checking card visibility
	print("Checking card visibility...")
	assert.False(card.isVisible)

	// Checking card "Click"
	print("Clicking on card...")
	card.Click()
	// Checking card visibility
	print("Checking card visibility...")
	assert.True(card.isVisible)

	print("End of card test!")

	// Missing: Display tests passed
}
