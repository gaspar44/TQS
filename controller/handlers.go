package controller

import (
	"encoding/json"
	"fmt"
	"gaspar44/TQS/model"
	"net/http"
)

const (
	Port       string = ":8080"
	fileSystem        = http.Dir("./view")
)

var (
	mux         map[string]func(writer http.ResponseWriter, requestBody *http.Request)
	activeGames map[string]*model.Game
)

type defaultHandler struct{}

func (handler *defaultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Print(request.URL.Path) // TODO: check on HTTP parameters
	if handlerFunction, ok := mux[request.URL.String()]; ok {
		handlerFunction(writer, request)
	}
}

func createGame(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
		return
	}

	var bodyData requestData
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&bodyData); err != nil {
		http.Error(writer, "Error on decode JSON", http.StatusBadRequest)
		return
	}

	playerName := bodyData.PlayerName
	gameDifficulty := model.Difficulty(bodyData.GameDifficulty)

	if _, exists := activeGames[playerName]; exists {
		writer.WriteHeader(http.StatusOK)
		return
	}

	game, err := model.NewGame(playerName, gameDifficulty)
	if err != nil {
		http.Error(writer, "Error while creating game", http.StatusInternalServerError)
		return
	}

	activeGames[playerName] = game
	writer.WriteHeader(http.StatusCreated)
}

func NewServer() *http.Server {
	server := http.Server{
		Addr:    Port,
		Handler: &defaultHandler{},
	}

	activeGames = make(map[string]*model.Game)
	mux = make(map[string]func(writer http.ResponseWriter, requestBody *http.Request))

	fileServer := http.FileServer(fileSystem)
	mux["/"] = fileServer.ServeHTTP
	mux["createGame"] = createGame

	return &server
}
