package model

type Player struct {
	Name string `json:"player_name,omitempty"`
	Time int    `json:"time,omitempty"`
}
