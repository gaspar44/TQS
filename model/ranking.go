package model

import (
	"gaspar44/TQS/model/custom_errors"
	"sync"
)

var (
	lock     = &sync.Mutex{}
	instance *Ranking
)

type Ranking struct {
	Players       *[]Player
	isInitialized bool
}

func (r *Ranking) GetPlayers() (*[]Player, error) {
	if r.Players == nil {
		return nil, custom_errors.NewRankingInitializationErrorWithMessage()
	}

	return r.Players, nil
}

func (r *Ranking) SetPlayers(players *[]Player) {
	if !instance.isInitialized {
		lock.Lock()
		defer lock.Unlock()

		if !instance.isInitialized {
			instance.Players = players
			instance.isInitialized = true
		}
	}
}

func (r *Ranking) release() {
	if instance != nil {
		lock.Lock()
		defer lock.Unlock()

		if instance != nil {
			instance = nil
		}
	}
}

func GetRankingInstance() (*Ranking, error) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Ranking{isInitialized: false,
				Players: nil}
		}
	}
	return instance, nil
}
