package model

type Card struct {
	isVisible bool
	value     int
}

func NewCard(assignedValue int) *Card {
	return &Card{
		isVisible: false,
		value:     assignedValue,
	}
}

func (c *Card) GetVisibility() bool {
	return c.isVisible
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) Click() {
	c.isVisible = !c.isVisible
}
