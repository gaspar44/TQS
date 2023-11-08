package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"gaspar44/TQS/model"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	assert2 "github.com/stretchr/testify/assert"
)

const (
	serverUrl     = "http://localhost:8080"
	debugFileName = "requests.txt"
)

func TestMain(m *testing.M) {
	srv := setup()
	code := m.Run()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer func() {
		cancel()
	}()

	srv.Shutdown(ctx)
	os.Exit(code)
}

func setup() *http.Server {
	logFile, err := setupDebugFile()
	if err != nil {
		panic(err)
	}
	srv := NewServerWithLogger(logFile)
	go func() {
		srv.ListenAndServe()
	}()

	return srv
}

func setupDebugFile() (io.Writer, error) {
	if _, err := os.Stat(debugFileName); err == nil {
		er := os.Remove(debugFileName)
		if er != nil {
			panic(err)
		}
	}

	file, err := os.Create(debugFileName)
	return file, err

}

func TestCreateGame(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)
	assert.Nil(err)
	assert.NotEmpty(body)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	var gameResponse createGameResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&gameResponse)

	assert.Nil(err)
	assert.Equal(http.StatusCreated, response.StatusCode)
	assert.Equal(playerName, gameResponse.PlayerName)
	assert.Equal(model.EasyDifficultyCardsTotal, len(gameResponse.Cards))
}

func TestCreateGameTwice(t *testing.T) {
	assert := assert2.New(t)
	playerName := "testTwice"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)
	assert.Nil(err)
	assert.NotEmpty(body)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	var gameResponse createGameResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&gameResponse)

	assert.Nil(err)
	assert.Equal(http.StatusCreated, response.StatusCode)
	assert.Equal(playerName, gameResponse.PlayerName)
	assert.Equal(model.EasyDifficultyCardsTotal, len(gameResponse.Cards))

	newResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)
	assert.Equal(http.StatusOK, newResponse.StatusCode)

	var newGameResponse createGameResponse
	err = json.NewDecoder(newResponse.Body).Decode(&newGameResponse)

	assert.Nil(err)
	assert.Equal(playerName, newGameResponse.PlayerName)
	assert.Equal(gameResponse.Cards, newGameResponse.Cards)
}

func TestCreateGameInvalidRequest(t *testing.T) {
	assert := assert2.New(t)

	type badRequest struct {
		Hello   int  `json:"hello,omitempty"`
		Success bool `json:"success,omitempty"`
	}
	request := badRequest{
		Hello:   10,
		Success: true,
	}

	body, err := json.Marshal(request)
	assert.Nil(err)
	assert.NotEmpty(body)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, response.StatusCode)
}

func TestCreateGameUnsupportedType(t *testing.T) {
	assert := assert2.New(t)

	type badRequest struct {
		Hello   int  `xml:"hello,omitempty"`
		Success bool `xml:"success,omitempty"`
	}
	request := badRequest{
		Hello:   10,
		Success: true,
	}

	body, err := xml.Marshal(request)
	assert.Nil(err)
	assert.NotEmpty(body)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	assert.Nil(err)
	assert.Equal(http.StatusUnsupportedMediaType, response.StatusCode)
}

func TestCreateGameGetMethod(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)
	assert.Nil(err)
	assert.NotEmpty(body)

	gameCreationRequest, err := http.NewRequest(http.MethodGet, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	assert.Nil(err)
	assert.Equal(http.StatusMethodNotAllowed, response.StatusCode)
	assert.Equal(http.MethodPost, response.Header.Get("Access-Control-Allow-Methods"))
}
