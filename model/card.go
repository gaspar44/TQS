package model

type Card struct {
	isVisible bool
	image     []byte
	value     int
}

func NewCard() *Card {
	return &Card{
		isVisible: false,
		image:     nil,
		value:     0,
	}
}

func (c *Card) GetVisibility() bool {
	return c.isVisible
}

func (c *Card) GetImage() []byte {
	return c.image
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) Click() {
	c.isVisible = !c.isVisible
}
