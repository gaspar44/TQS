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

	assert.Equal("application/json", response.Header.Get("Content-Type"))

	var gameResponse createGameResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&gameResponse)

	assert.Nil(err)
	assert.Equal(http.StatusCreated, response.StatusCode)
	assert.Equal(playerName, gameResponse.PlayerName)
	assert.Equal(0, gameResponse.Points)
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
	assert.Equal("application/json", response.Header.Get("Content-Type"))

	var gameResponse createGameResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&gameResponse)

	assert.Nil(err)
	assert.Equal(http.StatusCreated, response.StatusCode)
	assert.Equal(0, gameResponse.Points)
	assert.Equal(playerName, gameResponse.PlayerName)
	assert.Equal(model.EasyDifficultyCardsTotal, len(gameResponse.Cards))

	newResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)
	assert.Equal(http.StatusOK, newResponse.StatusCode)

	var newGameResponse createGameResponse
	err = json.NewDecoder(newResponse.Body).Decode(&newGameResponse)

	assert.Nil(err)
	assert.Equal(playerName, newGameResponse.PlayerName)
	assert.Equal(gameResponse.Points, newGameResponse.Points)
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

func TestCreateGameOptionsMethod(t *testing.T) {
	assert := assert2.New(t)

	gameCreationRequest, err := http.NewRequest(http.MethodOptions, serverUrl+CreateGame, nil)
	assert.Nil(err)

	gameCreationRequest.Header.Set("Access-Control-Request-Headers", "content-type")
	gameCreationRequest.Header.Set("Access-Control-Request-Method", http.MethodPost)
	gameCreationRequest.Header.Set("Sec-Fetch-Mode", "cors")
	gameCreationRequest.Header.Set("Sec-Fetch-Site", "same-site")

	client := &http.Client{}
	response, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	defer response.Body.Close()

	assert.Nil(err)
	assert.Equal(http.StatusOK, response.StatusCode)
	assert.Equal(http.MethodGet+" ,"+http.MethodPost, response.Header.Get("Access-Control-Allow-Methods"))
	assert.Equal("Content-Type", response.Header.Get("Access-Control-Allow-Headers"))
	assert.Equal("*", response.Header.Get("Access-Control-Allow-Origin"))
}

func TestDefaultLoadingPage(t *testing.T) {
	assert := assert2.New(t)
	defaultPageRequest, err := http.NewRequest(http.MethodGet, serverUrl, nil)
	assert.Nil(err)

	client := &http.Client{}
	response, err := client.Do(defaultPageRequest)

	assert.Nil(err)
	message, err := io.ReadAll(response.Body)
	assert.Nil(err)

	defer response.Body.Close()

	assert.Equal(http.StatusOK, response.StatusCode)
	assert.Equal(WelcomeMessage, string(message))
}

func TestChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "choose card"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	createGameHttpResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	decoderGame := json.NewDecoder(createGameHttpResponse.Body)
	defer createGameHttpResponse.Body.Close()

	var createdGame createGameResponse
	err = decoderGame.Decode(&createdGame)

	assert.Nil(err)

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: 0,
	}

	chooseCardRequestBody, err := json.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodPost, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	choseCardRequest.Header.Set("Content-Type", "application/json")
	assert.Nil(err)

	choseCardHttpResponse, err := client.Do(choseCardRequest)

	assert.Nil(err)
	assert.Equal(http.StatusOK, choseCardHttpResponse.StatusCode)
	assert.Equal("application/json", choseCardHttpResponse.Header.Get("Content-Type"))

	decoder := json.NewDecoder(choseCardHttpResponse.Body)

	defer choseCardHttpResponse.Body.Close()

	var cardChoice choiceCardResponse
	err = decoder.Decode(&cardChoice)

	assert.Nil(err)
	assert.Equal(len(createdGame.Cards), len(cardChoice.Cards))
	assert.Equal(1, cardChoice.Points)
	assert.NotEqual(createdGame.Cards[0].IsVisible, cardChoice.Cards[0].IsVisible)

}

func TestChooseCardUnsupportedMedia(t *testing.T) {
	assert := assert2.New(t)
	playerName := "choose card"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	createGameHttpResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	decoderGame := json.NewDecoder(createGameHttpResponse.Body)
	defer createGameHttpResponse.Body.Close()

	var createdGame createGameResponse
	err = decoderGame.Decode(&createdGame)

	assert.Nil(err)

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: 0,
	}

	chooseCardRequestBody, err := xml.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodPost, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	assert.Nil(err)

	choseCardHttpResponse, err := client.Do(choseCardRequest)

	assert.Nil(err)
	assert.Equal(http.StatusUnsupportedMediaType, choseCardHttpResponse.StatusCode)
}

// Partition
func TestWrongChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "choose wrong card"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	createGameHttpResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	decoderGame := json.NewDecoder(createGameHttpResponse.Body)
	defer createGameHttpResponse.Body.Close()

	var createdGame createGameResponse
	err = decoderGame.Decode(&createdGame)

	assert.Nil(err)

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: model.EasyDifficultyCardsTotal + 1,
	}

	chooseCardRequestBody, err := json.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodPost, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	assert.Nil(err)
	choseCardRequest.Header.Set("Content-Type", "application/json")

	choseCardHttpResponse, err := client.Do(choseCardRequest)

	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, choseCardHttpResponse.StatusCode)
}

// Partition
func TestChooseInvalidCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "choose wrong card"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	createGameHttpResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	decoderGame := json.NewDecoder(createGameHttpResponse.Body)
	defer createGameHttpResponse.Body.Close()

	var createdGame createGameResponse
	err = decoderGame.Decode(&createdGame)

	assert.Nil(err)

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: -1,
	}

	chooseCardRequestBody, err := json.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodPost, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	assert.Nil(err)
	choseCardRequest.Header.Set("Content-Type", "application/json")

	choseCardHttpResponse, err := client.Do(choseCardRequest)

	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, choseCardHttpResponse.StatusCode)
}

func TestChooseCardMethodNotAllowed(t *testing.T) {
	assert := assert2.New(t)
	playerName := "choose card"

	request := createGameRequest{
		PlayerName:     playerName,
		GameDifficulty: model.Easy,
	}

	body, err := json.Marshal(request)

	gameCreationRequest, err := http.NewRequest(http.MethodPost, serverUrl+CreateGame, bytes.NewBuffer(body))
	assert.Nil(err)
	gameCreationRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	createGameHttpResponse, err := client.Do(gameCreationRequest)
	assert.Nil(err)

	decoderGame := json.NewDecoder(createGameHttpResponse.Body)
	defer createGameHttpResponse.Body.Close()

	var createdGame createGameResponse
	err = decoderGame.Decode(&createdGame)

	assert.Nil(err)

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: 0,
	}

	chooseCardRequestBody, err := json.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodGet, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	assert.Nil(err)
	choseCardRequest.Header.Set("Content-Type", "application/json")

	choseCardHttpResponse, err := client.Do(choseCardRequest)

	assert.Nil(err)
	assert.Equal(http.StatusMethodNotAllowed, choseCardHttpResponse.StatusCode)
	assert.Equal(http.MethodPost, choseCardHttpResponse.Header.Get("Access-Control-Allow-Methods"))

}

func TestChooseCardMethodUncreatedGame(t *testing.T) {
	assert := assert2.New(t)
	playerName := "not created player"

	chooseCardJsonRequest := choiceCardRequest{
		PlayerName: playerName,
		CardChoice: 0,
	}

	chooseCardRequestBody, err := json.Marshal(chooseCardJsonRequest)
	assert.Nil(err)
	assert.NotEmpty(chooseCardRequestBody)

	choseCardRequest, err := http.NewRequest(http.MethodPost, serverUrl+ChooseCard, bytes.NewBuffer(chooseCardRequestBody))
	assert.Nil(err)
	choseCardRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	choseCardHttpResponse, err := client.Do(choseCardRequest)
	assert.Nil(err)

	defer choseCardHttpResponse.Body.Close()
	assert.Equal(http.StatusNotFound, choseCardHttpResponse.StatusCode)
}

func TestGetRanking(t *testing.T) {
	assert := assert2.New(t)

	getRankingRequest, err := http.NewRequest(http.MethodGet, serverUrl+GetRanking, nil)
	assert.Nil(err)

	client := &http.Client{}
	getRankingResponse, err := client.Do(getRankingRequest)
	assert.Nil(err)
	assert.Equal(http.StatusOK, getRankingResponse.StatusCode)
	assert.Equal("application/json", getRankingResponse.Header.Get("Content-Type"))

	defer getRankingResponse.Body.Close()

	var rankingResponse model.Ranking
	decoder := json.NewDecoder(getRankingResponse.Body)
	err = decoder.Decode(&rankingResponse)
	assert.Nil(err)
}

func TestGetRankingPostMethod(t *testing.T) {
	assert := assert2.New(t)

	getRankingRequest, err := http.NewRequest(http.MethodPost, serverUrl+GetRanking, nil)
	assert.Nil(err)

	client := &http.Client{}
	getRankingResponse, err := client.Do(getRankingRequest)
	assert.Nil(err)

	defer getRankingResponse.Body.Close()

	assert.Equal(http.StatusMethodNotAllowed, getRankingResponse.StatusCode)
	assert.Equal(http.MethodGet, getRankingResponse.Header.Get("Access-Control-Allow-Methods"))
}

func TestEndGame(t *testing.T) {
	assert := assert2.New(t)
	playerName := "game over"

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

	endRequest := endGameRequest{
		PlayerName: playerName,
	}

	endRequestBody, err := json.Marshal(endRequest)
	assert.Nil(err)
	assert.NotEmpty(endRequestBody)

	endGameHttpRequest, err := http.NewRequest(http.MethodPost, serverUrl+EndGame, bytes.NewBuffer(endRequestBody))
	assert.Nil(err)
	endGameHttpRequest.Header.Set("Content-Type", "application/json")

	gameEndResponse, err := client.Do(endGameHttpRequest)
	assert.Nil(err)
	defer gameEndResponse.Body.Close()

	assert.Equal(http.StatusOK, gameEndResponse.StatusCode)
	assert.Equal("application/json", gameEndResponse.Header.Get("Content-Type"))

	var endResponse endGameResponse
	decoder := json.NewDecoder(gameEndResponse.Body)
	decoder.Decode(&endResponse)
	assert.Equal(endRequest.PlayerName, endResponse.PlayerName)
}

func TestEndGameWrongMethod(t *testing.T) {
	assert := assert2.New(t)
	playerName := "game over"

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

	endRequest := endGameRequest{
		PlayerName: playerName,
	}

	endRequestBody, err := json.Marshal(endRequest)
	assert.Nil(err)
	assert.NotEmpty(endRequestBody)

	endGameHttpRequest, err := http.NewRequest(http.MethodGet, serverUrl+EndGame, bytes.NewBuffer(endRequestBody))
	assert.Nil(err)
	endGameHttpRequest.Header.Set("Content-Type", "application/json")

	gameEndResponse, err := client.Do(endGameHttpRequest)
	assert.Nil(err)
	defer gameEndResponse.Body.Close()

	assert.Equal(http.StatusMethodNotAllowed, gameEndResponse.StatusCode)
	assert.Equal(http.MethodPost, gameEndResponse.Header.Get("Access-Control-Allow-Methods"))
}

func TestEndNotCreatedGame(t *testing.T) {
	assert := assert2.New(t)
	playerName := "game over!"

	endRequest := endGameRequest{
		PlayerName: playerName,
	}

	endRequestBody, err := json.Marshal(endRequest)
	assert.Nil(err)
	assert.NotEmpty(endRequestBody)

	endGameHttpRequest, err := http.NewRequest(http.MethodPost, serverUrl+EndGame, bytes.NewBuffer(endRequestBody))
	assert.Nil(err)

	endGameHttpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	gameEndResponse, err := client.Do(endGameHttpRequest)
	assert.Nil(err)
	defer gameEndResponse.Body.Close()

	assert.Equal(http.StatusNotFound, gameEndResponse.StatusCode)
}

func TestEndGameUnsupportedMedia(t *testing.T) {
	assert := assert2.New(t)
	playerName := "game over!"

	endRequest := endGameRequest{
		PlayerName: playerName,
	}

	endRequestBody, err := json.Marshal(endRequest)
	assert.Nil(err)
	assert.NotEmpty(endRequestBody)

	endGameHttpRequest, err := http.NewRequest(http.MethodPost, serverUrl+EndGame, bytes.NewBuffer(endRequestBody))
	assert.Nil(err)

	endGameHttpRequest.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	gameEndResponse, err := client.Do(endGameHttpRequest)
	assert.Nil(err)
	defer gameEndResponse.Body.Close()

	assert.Equal(http.StatusUnsupportedMediaType, gameEndResponse.StatusCode)
}

func TestEndGameBadRequest(t *testing.T) {
	assert := assert2.New(t)
	playerName := "game over!"

	endRequest := endGameRequest{
		PlayerName: playerName,
	}

	endRequestBody, err := xml.Marshal(endRequest)
	assert.Nil(err)
	assert.NotEmpty(endRequestBody)

	endGameHttpRequest, err := http.NewRequest(http.MethodPost, serverUrl+EndGame, bytes.NewBuffer(endRequestBody))
	assert.Nil(err)

	endGameHttpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	gameEndResponse, err := client.Do(endGameHttpRequest)
	assert.Nil(err)
	defer gameEndResponse.Body.Close()

	assert.Equal(http.StatusBadRequest, gameEndResponse.StatusCode)
}
