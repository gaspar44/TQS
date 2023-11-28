package storage

import (
	"encoding/json"
	"gaspar44/TQS/model"
	"os"
)

const (
	defaultStorageName = "ranking.json"
)

type JsonStorage struct {
	storageFileName string
}

func NewJsonStorage(fileName string) JsonStorage {
	return JsonStorage{
		storageFileName: fileName,
	}
}

func NewDefaultJsonStorage() JsonStorage {
	return NewJsonStorage(defaultStorageName)
}

func (jsStorage JsonStorage) ReadRanking() (*model.Ranking, error) {
	if _, statErr := os.Stat(jsStorage.storageFileName); statErr == nil {
		rawData, err := os.ReadFile(jsStorage.storageFileName)

		if err != nil {
			return nil, err
		}

		ranking := model.GetRankingInstance()
		var players []model.Player
		err = json.Unmarshal(rawData, &players)

		if err != nil {
			players = make(model.Players, 0)
			ranking.SetPlayers(players)
			return ranking, nil
		}

		ranking.SetPlayers(players)
		return ranking, nil
	}

	return model.GetRankingInstance(), nil
}

func (jsStorage JsonStorage) WriteRanking(ranking *model.Ranking) error {
	fileInfo, err := os.Create(jsStorage.storageFileName)

	if err != nil {
		return err
	}

	defer fileInfo.Close()

	players, err := ranking.GetPlayers()

	if err != nil {
		return err
	}

	rawData, err := json.Marshal(players)

	fileInfo.Write(rawData)

	return nil
}
