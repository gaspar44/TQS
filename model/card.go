package model

type Card struct {
	// Info about the card (visibility and value)
	isVisible bool
	isDisable bool
	value     int
}

// Constructor
func NewCard(assignedValue int) Card {
	return Card{
		isVisible: false,
		isDisable: false,
		value:     assignedValue,
	}
}

// Getters
func (c *Card) GetVisibility() bool {
	// To get card visibility
	return c.isVisible
}

func (c *Card) GetValue() int {
	// Common getter
	return c.value
}

// Setters
func (c *Card) SetVisibility(assignedVisibility bool) {
	c.isVisible = assignedVisibility
}

func (c *Card) SetValue(assignedValue int) {
	c.value = assignedValue
}

func (c *Card) disable() {
	c.isDisable = true
}

// Functions:

func (c *Card) Click() {
	// Missing:
	// Check if card is already matched with another
	if c.isVisible != true {
		c.isVisible = !c.isVisible
	}
}
