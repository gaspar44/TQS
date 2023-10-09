package model

import (
	"fmt"
	"gaspar44/TQS/model/errors"
	"gaspar44/TQS/storage"
	"sync"
)

type Ranking struct {
	Players *[]Player
	Storage storage.StorageHandler
}

func (r *Ranking) Write() error {
	return r.Storage.WriteRanking()
}

var (
	lock     = &sync.Mutex{}
	instance *Ranking
)

func GetRankingInstance(storage storage.StorageHandler) (*Ranking, error) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			readBytes, err := storage.ReadRanking()

			if err != nil {
				return nil, errors.NewRankingInitializationError(err)
			}

			players := convertToPlayers(readBytes)
			instance = &Ranking{Storage: storage,
				Players: players}
		}
	}

	return instance, nil
}

func convertToPlayers(readBytes []byte) *[]Player {
	// TODO
	fmt.Print(readBytes)
	return nil
}
