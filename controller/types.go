package controller

import "gaspar44/TQS/model"

const (
	WelcomeMessage = "Welcome to the memory game!. Integration will be in other step"
)

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

type endGameRequest struct {
	PlayerName string `json:"player_name"`
}

type endGameResponse struct {
	PlayerName string `json:"player_name"`
	Points     int    `json:"points"`
}
