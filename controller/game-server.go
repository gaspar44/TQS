package controller

import (
	"gaspar44/TQS/model"
	"gaspar44/TQS/model/storage"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	Port string = ":8080"
)

var (
	infoLogger     *log.Logger
	debugLogger    *log.Logger
	ranking        *model.Ranking
	rankingStorage storage.StorageHandler
)

func NewServer() *http.Server {
	return NewServerWithLogger(os.Stdout)
}

func NewServerWithLogger(out io.Writer) *http.Server {
	infoLogger = log.New(out, "INFO: ", log.LstdFlags|log.Llongfile)
	debugLogger = log.New(out, "DEBUG: ", log.LstdFlags|log.Llongfile)

	server := http.Server{
		Addr:    Port,
		Handler: &defaultHandler{},
	}

	activeGames = make(map[string]*model.Game)
	mux = make(map[string]func(writer http.ResponseWriter, requestBody *http.Request))

	mux["/"] = welcome
	mux[CreateGame] = createGame
	mux[GetRanking] = displayRanking
	mux[ChooseCard] = chooseCard
	mux[EndGame] = endGame

	infoLogger.Println("Server created and running at port: " + server.Addr)
	rankingStorage = storage.StorageHandler(storage.NewDefaultJsonStorage())
	ranking, _ = rankingStorage.ReadRanking()
	infoLogger.Println("Ranking created")
	return &server
}
