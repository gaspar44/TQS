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

// Getters
func (c *Card) GetVisibility() bool {
	return c.isVisible
}

func (c *Card) GetValue() int {
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
	if c.isVisible == false && c.isDisable == true {
		c.isVisible = !c.isVisible
	}
}
