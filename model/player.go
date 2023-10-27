package model

type Player struct {
	Name string `json:"player_name,omitempty"`
	Time int64  `json:"time,omitempty"`
}

type Players []Player

// Implementation of sort.Interface

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i int, j int) bool {
	return p[i].Time < p[j].Time
}

func (p Players) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
