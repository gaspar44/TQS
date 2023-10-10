package model

import (
	"gaspar44/TQS/model/custom_errors"
	assert2 "github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetRankingInstance(t *testing.T) {
	assert := assert2.New(t)
	ranking, err := GetRankingInstance()
	defer ranking.release()

	assert.NotNil(ranking)
	assert.Nil(err)
	assert.False(ranking.isInitialized)

	ranking2, err := GetRankingInstance()
	assert.Nil(err)
	assert.Equal(ranking, ranking2)
}

func TestInitializeRanking(t *testing.T) {
	assert := assert2.New(t)

	ranking, err := GetRankingInstance()
	defer ranking.release()
	assert.Nil(err)
	assert.NotNil(ranking)
	assert.False(ranking.isInitialized)

	players := []Player{
		{
			Name: "test1",
			Time: 0,
		},
		{
			Name: "test2",
			Time: 10,
		},
	}

	ranking.SetPlayers(&players)
	assert.True(ranking.isInitialized)

	playersInRanking, err := ranking.GetPlayers()

	assert.Nil(err)
	assert.Equal(playersInRanking, &players)
}

func TestRankingInitializationError(t *testing.T) {
	assert := assert2.New(t)

	ranking, err := GetRankingInstance()
	defer ranking.release()
	assert.Nil(err)
	assert.NotNil(ranking)
	assert.False(ranking.isInitialized)

	players, err := ranking.GetPlayers()
	assert.Nil(players)

	if assert.NotNil(err) {
		assert.Equal(custom_errors.RankingInitializationErrorMessage, err.Error())
	}
}

func TestInitializeRankingMultiplesPlayers(t *testing.T) {
	assert := assert2.New(t)

	ranking, err := GetRankingInstance()
	defer ranking.release()
	assert.Nil(err)
	assert.NotNil(ranking)
	assert.False(ranking.isInitialized)

	players := []Player{
		{
			Name: "test1",
			Time: 0,
		},
		{
			Name: "test2",
			Time: 10,
		},
	}

	differentPlayersInitialization := []Player{
		{
			Name: "test3",
			Time: 10,
		},
	}

	ranking.SetPlayers(&players)
	assert.True(ranking.isInitialized)
	ranking.SetPlayers(&differentPlayersInitialization)
	playersInRanking, err := ranking.GetPlayers()

	assert.Nil(err)
	assert.Equal(playersInRanking, &players)
	assert.NotEqual(playersInRanking, &differentPlayersInitialization)
}

func TestGetRankingInstanceMultithreading(t *testing.T) {
	assert := assert2.New(t)
	var rankings [2]*Ranking
	var wg sync.WaitGroup

	for i := 0; i < len(rankings); i++ {
		wg.Add(1)
		go func(position int) {

			defer wg.Done()
			ranking, err := GetRankingInstance()
			assert.Nil(err)
			assert.NotNil(ranking)
			rankings[position] = ranking
		}(i)
	}

	wg.Wait()
	assert.Equal(rankings[0], rankings[1])
}
