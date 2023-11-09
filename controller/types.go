package controller

import "gaspar44/TQS/model"

type createGameRequest struct {
	PlayerName     string           `json:"player_name"`
	GameDifficulty model.Difficulty `json:"difficulty"`
}

type createGameResponse struct {
	PlayerName string       `json:"player_name"`
	Cards      []model.Card `json:"cards"`
	Points     int          `json:"points"`
}

type choiceCardRequest struct {
	PlayerName string `json:"player_name"`
	CardChoice int    `json:"card_choice"`
}

type choiceCardResponse struct {
	PlayerName string       `json:"player_name"`
	Success    bool         `json:"success"`
	Cards      []model.Card `json:"cards"`
	Points     int          `json:"points"`
}
