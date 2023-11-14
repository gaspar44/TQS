package controller

import (
	"encoding/json"
	"fmt"
	"gaspar44/TQS/model"
	"net/http"
	"net/http/httputil"
)

var (
	mux         map[string]func(writer http.ResponseWriter, requestBody *http.Request)
	activeGames map[string]*model.Game
)

type defaultHandler struct{}

func (handler *defaultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Print(request.URL.Path) // TODO: check on HTTP parameters
	if handlerFunction, ok := mux[request.URL.String()]; ok {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		if request.Method == http.MethodOptions {
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			writer.Header().Set("Access-Control-Allow-Methods", http.MethodGet+" ,"+http.MethodPost)
			return
		}

		handlerFunction(writer, request)
	}
}
func welcome(writer http.ResponseWriter, request *http.Request) {
	dumpHttpRequest, _ := httputil.DumpRequest(request, true)
	debugLogger.Println(string(dumpHttpRequest))
	defer request.Body.Close()

	writer.Write([]byte("Welcome to the memory game!. Integration will be in other step"))
}

func createGame(writer http.ResponseWriter, request *http.Request) {
	dumpHttpRequest, _ := httputil.DumpRequest(request, true)
	debugLogger.Println(string(dumpHttpRequest))
	defer request.Body.Close()
	if request.Method != http.MethodPost {
		infoLogger.Println("Invalid http method:" + request.Method)
		writer.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if request.Header.Get("Content-Type") != "application/json" {
		infoLogger.Println("Unsupported type:" + request.Header.Get("Content-Type"))
		writer.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var bodyData createGameRequest
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&bodyData); err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Error on decode JSON", http.StatusBadRequest)
		return
	}

	playerName := bodyData.PlayerName
	gameDifficulty := bodyData.GameDifficulty

	if _, exists := activeGames[playerName]; exists {
		infoLogger.Println("Game already exists")
		game := activeGames[playerName]
		cards := game.GetCards()
		points := game.GetPoints()

		response := createGameResponse{
			PlayerName: playerName,
			Cards:      cards,
			Points:     points,
		}

		writer.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(response)
		if err != nil {
			debugLogger.Println(err.Error())
		}
		return
	}

	game, err := model.NewGame(playerName, gameDifficulty)

	if err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Error while creating game", http.StatusInternalServerError)
		return
	}

	activeGames[playerName] = game
	infoLogger.Println("Game created")

	cards := game.GetCards()
	response := createGameResponse{
		PlayerName: playerName,
		Cards:      cards,
	}

	encoder := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = encoder.Encode(response)

	if err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Error on JSON marshaling", http.StatusBadRequest)
		return
	}
}

func chooseCard(writer http.ResponseWriter, request *http.Request) {
	dumpHttpRequest, _ := httputil.DumpRequest(request, true)
	debugLogger.Println(string(dumpHttpRequest))

	if request.Method != http.MethodPost {
		infoLogger.Println("Invalid http method:" + request.Method)
		writer.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var choiceRequest choiceCardRequest
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&choiceRequest); err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Error on decode JSON", http.StatusBadRequest)
		return
	}

	if _, ok := activeGames[choiceRequest.PlayerName]; !ok {
		infoLogger.Println("Game for player " + choiceRequest.PlayerName + " has not been created.")
		http.Error(writer, "No active games for "+choiceRequest.PlayerName, http.StatusNotFound)
		return
	}
	game := activeGames[choiceRequest.PlayerName]

	correct, err := game.ChooseCardOnBoard(choiceRequest.CardChoice)

	if err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Invalid card", http.StatusBadRequest)
		return
	}

	choiceResponse := choiceCardResponse{
		PlayerName: choiceRequest.PlayerName,
		Success:    correct,
		Cards:      game.GetCards(),
		Points:     game.GetPoints(),
	}

	encoder := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	err = encoder.Encode(choiceResponse)

	if err != nil {
		debugLogger.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func displayRanking(writer http.ResponseWriter, request *http.Request) {
	dumpHttpRequest, _ := httputil.DumpRequest(request, true)
	debugLogger.Println(string(dumpHttpRequest))

	if request.Method != http.MethodGet {
		infoLogger.Println("Invalid http method:" + request.Method)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
		return
	}

	ranking, err := model.GetRankingInstance()

	if err != nil {
		debugLogger.Println(err.Error())
		http.Error(writer, "Error while getting ranking", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	err = encoder.Encode(ranking)

	if err != nil {
		http.Error(writer, "Error on JSON ranking", http.StatusInternalServerError)
	}

}
