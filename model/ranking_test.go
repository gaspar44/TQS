package model

import (
	"gaspar44/TQS/model/custom_errors"
	assert2 "github.com/stretchr/testify/assert"
	"strconv"
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
			Name:   "test1",
			Points: 0,
		},
		{
			Name:   "test2",
			Points: 10,
		},
	}

	ranking.SetPlayers(players)
	assert.True(ranking.isInitialized)

	playersInRanking, err := ranking.GetPlayers()

	assert.Nil(err)
	assert.Equal(playersInRanking, players)
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
			Name:   "test1",
			Points: 0,
		},
		{
			Name:   "test2",
			Points: 10,
		},
	}

	differentPlayersInitialization := []Player{
		{
			Name:   "test3",
			Points: 10,
		},
	}

	ranking.SetPlayers(players)
	assert.True(ranking.isInitialized)
	ranking.SetPlayers(differentPlayersInitialization)
	playersInRanking, err := ranking.GetPlayers()

	assert.Nil(err)
	assert.Equal(playersInRanking, players)
	assert.NotEqual(playersInRanking, differentPlayersInitialization)
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

func TestAddPlayersToRankingAtMax(t *testing.T) {
	assert := assert2.New(t)
	players := make(Players, maxPlayers)

	for i := 0; i < maxPlayers; i++ {
		newPlayer := Player{
			Name:   "test " + strconv.Itoa(i),
			Points: i + 1,
		}
		players[i] = newPlayer
	}
	ranking, err := GetRankingInstance()
	assert.NotNil(ranking)
	assert.Nil(err)
	defer ranking.release()

	ranking.SetPlayers(players)
	newPlayerRecord := Player{
		Name:   "newRecord",
		Points: 0, // Virtually impossible
	}

	ranking.Update(newPlayerRecord)
	newRanking, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.Equal(maxPlayers, len(newRanking))
	assert.Equal(newPlayerRecord, newRanking[0])
}

func TestAddPlayersEmptyRanking(t *testing.T) {
	assert := assert2.New(t)

	for i := 0; i < maxPlayers; i++ {
		ranking, err := GetRankingInstance()
		assert.Nil(err)
		assert.NotNil(ranking)

		if i == 0 { // To avoid collision with other test. But defer must be called once.
			defer ranking.release()
		}

		assert.Equal(i, len(ranking.Players))

		newPlayer := Player{
			Name:   "test " + strconv.Itoa(i),
			Points: i + 1,
		}

		ranking.Update(newPlayer)
	}
}

func TestAddPlayersToRankingWithoutDeserve(t *testing.T) {
	assert := assert2.New(t)
	players := make(Players, maxPlayers)

	for i := 0; i < maxPlayers; i++ {
		newPlayer := Player{
			Name:   "test " + strconv.Itoa(i),
			Points: i + 1,
		}
		players[i] = newPlayer
	}
	ranking, err := GetRankingInstance()
	assert.NotNil(ranking)
	assert.Nil(err)

	ranking.SetPlayers(players)
	newPlayerRecord := Player{
		Name:   "newRecord",
		Points: 10000,
	}

	ranking.Update(newPlayerRecord)
	newRanking, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.Equal(maxPlayers, len(newRanking))

	for player := range newRanking {
		assert.NotEqual(player, newPlayerRecord)
	}
}
