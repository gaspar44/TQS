package model

import (
	"gaspar44/TQS/model/custom_errors"
	"sort"
	"sync"
)

const (
	maxPlayers = 10
)

var (
	lock     = &sync.Mutex{}
	instance *Ranking
)

type Ranking struct {
	Players       []Player
	isInitialized bool
}

func (r *Ranking) GetPlayers() ([]Player, error) {
	if r.Players == nil {
		return nil, custom_errors.NewRankingInitializationErrorWithMessage()
	}

	return r.Players, nil
}

func (r *Ranking) SetPlayers(players []Player) {
	if !instance.isInitialized {
		lock.Lock()
		defer lock.Unlock()

		if !instance.isInitialized {
			instance.Players = players
			sort.Sort(Players(r.Players))
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

func (r *Ranking) Update(player Player) {
	if len(r.Players) < maxPlayers {
		r.Players = append(r.Players, player)
		sort.Sort(Players(r.Players))
		return
	}

	if r.Players[0].Points < player.Points {
		return // Player doesn't deserve to be in the top 10 ranking
	}

	r.Players = append(r.Players, player)
	sort.Sort(Players(r.Players))
	r.Players = r.Players[:len(r.Players)-1] // Remove the lowest rankings
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

//////////////////////////////////////////////////////////////////////////////////////
