package model

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

// Unit Test
func TestCard(t *testing.T) {
	assert := assert2.New(t)
	card := NewCard(1)

	assert.Equal(1, card.Value)
	assert.False(card.IsVisible)

	card.Click()
	assert.True(card.IsVisible)
}
