package model

type Player struct {
	Name   string `json:"player_name,omitempty"`
	Points int    `json:"points,omitempty"`
}

type Players []Player

// Implementation of sort.Interface

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i int, j int) bool {
	return p[i].Points < p[j].Points
}

func (p Players) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
