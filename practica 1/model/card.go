package model

type Card struct {
	IsVisible bool `json:"visible"`
	IsDisable bool `json:"disable"`
	Value     int  `json:"value"`
}

// Constructor
func NewCard(assignedValue int) Card {
	return Card{
		IsVisible: false,
		IsDisable: false,
		Value:     assignedValue,
	}
}

func (c *Card) GetValue() int {
	return c.Value
}

func (c *Card) disable() {
	c.IsDisable = true
}

// Functions:
func (c *Card) Click() {
	if !c.IsVisible && !c.IsDisable {
		c.IsVisible = !c.IsVisible
	}
}
