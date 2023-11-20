package storage

import (
	"encoding/json"
	"gaspar44/TQS/model"
	"gaspar44/TQS/model/custom_errors"
	"os"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	assert := assert2.New(t)
	ranking := model.GetRankingInstance()
	defer ranking.Release()
	players := []model.Player{
		{
			Name:   "test1",
			Points: 1,
		},
		{
			Name:   "test2",
			Points: 10,
		},
	}

	ranking.SetPlayers(players)
	handler := NewDefaultJsonStorage()
	err := handler.WriteRanking(ranking)
	assert.Nil(err)
	_, err = os.Stat(defaultStorageName)
	assert.Nil(err)
}

func TestRead(t *testing.T) {
	assert := assert2.New(t)
	expectedPlayers := createPlayersFromFile("ranking_test.txt")

	handler := NewJsonStorage("ranking_test.txt")
	ranking, err := handler.ReadRanking()

	assert.Nil(err)
	assert.NotNil(ranking)
	defer ranking.Release()

	players, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.NotEmpty(players)
	assert.Equal(expectedPlayers, model.Players(players))
}

func TestReadWrongFormatFile(t *testing.T) {
	assert := assert2.New(t)

	handler := NewJsonStorage("random_file.txt")
	ranking, err := handler.ReadRanking()

	assert.Nil(err)
	assert.NotNil(ranking)
	defer ranking.Release()

	players, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.Empty(players)
}

func TestWriteEmptyRanking(t *testing.T) {
	assert := assert2.New(t)

	handler := NewJsonStorage("/tmp/empty")
	ranking := model.GetRankingInstance()

	assert.NotNil(ranking)
	defer ranking.Release()

	err := handler.WriteRanking(ranking)
	assert.NotNil(err)
}

func TestReadFromNoExistingFile(t *testing.T) {
	assert := assert2.New(t)

	handler := NewJsonStorage("holi.txt")
	ranking, err := handler.ReadRanking()

	assert.Nil(err)
	assert.NotNil(ranking)
	defer ranking.Release()

	players, err := ranking.GetPlayers()
	assert.Nil(players)
	assert.NotNil(err)
	assert.Equal(custom_errors.RankingInitializationErrorMessage, err.Error())
}

func createPlayersFromFile(fileName string) model.Players {
	rawData, err := os.ReadFile(fileName)

	if err != nil {
		return model.Players{}
	}

	var players model.Players

	json.Unmarshal(rawData, &players)
	return players
}
