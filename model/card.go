package model

type Card struct {
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

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) disable() {
	c.isDisable = true
}

// Functions:
func (c *Card) Click() {
	if !c.isVisible && !c.isDisable {
		c.isVisible = !c.isVisible
	}
}
