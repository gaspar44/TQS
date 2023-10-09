package model

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestCardChangeVisibility(t *testing.T) {
	assert := assert2.New(t)
	card := NewCard(1)

	assert.Equal(1, card.value)
	assert.False(card.isVisible)
	card.Click()
	assert.True(card.isVisible)
}
