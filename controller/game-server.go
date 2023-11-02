package controller

import (
	"gaspar44/TQS/model"
	"log"
	"net/http"
	"os"
)

const (
	Port       string = ":8080"
	fileSystem        = http.Dir("./view")
)

var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Llongfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Llongfile)
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
	mux[CreateGame] = createGame
	mux[GetRanking] = displayRanking
	mux[GetCards] = displayCards
	mux[ChooseCard] = chooseCard

	infoLogger.Println("Server created and running at port: " + server.Addr)
	return &server
}
