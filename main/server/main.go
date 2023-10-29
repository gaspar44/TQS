package main

import (
	"encoding/json"
	"fmt"
	"gaspar44/TQS/model"
	"net/http"
)

func main() {
	// Main funcion to lunch game and server
	fmt.Print("Hello world")

	//start_game
	http.HandleFunc("/createGame", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var requestData struct {
				PlayerName     string `json:"playerName"`
				GameDifficulty int    `json:"gameDifficulty"`
			}

			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&requestData); err != nil {
				http.Error(w, "Error on decode JSON", http.StatusBadRequest)
				return
			}

			playerName := requestData.PlayerName
			gameDifficulty := model.Difficulty(requestData.GameDifficulty)

			game, err := model.NewGame(playerName, gameDifficulty)
			if err != nil {
				http.Error(w, "Error while creating game", http.StatusInternalServerError)
				return
			}

			game.Start()
			w.WriteHeader(http.StatusOK)
		} else {
			// Maybe GET (?)
			http.Error(w, "Wrong method...", http.StatusMethodNotAllowed)
		}
	})
}
