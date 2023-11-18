package model

import (
	"gaspar44/TQS/model/custom_errors"
	"strconv"
	"sync"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

// Unit Test: Checking "GetRankingInstance()" function
func TestGetRankingInstance(t *testing.T) {
	assert := assert2.New(t)
	ranking := GetRankingInstance()
	defer ranking.Release()

	assert.NotNil(ranking)

	assert.False(ranking.isInitialized)

	ranking2 := GetRankingInstance()
	assert.Equal(ranking, ranking2)
}

// Unit Test: Checking players added on ranking
// Partition Share: If 1 player is added correctly, any could be
func TestInitializeRanking(t *testing.T) {
	assert := assert2.New(t)

	ranking := GetRankingInstance()
	defer ranking.Release()
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

// Unit Test: Checking updating player
// Partition Share: If i can update a player, i can update any
func TestRankingInsertSamePlayerTwice(t *testing.T) {
	assert := assert2.New(t)

	ranking := GetRankingInstance()
	defer ranking.Release()
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

	ranking.Update(Player{
		Name:   "test1",
		Points: 0,
	})

	assert.Equal(len(players), len(ranking.Players))
}

// Unit Test: Checking error on initialization
// Limit values
func TestRankingInitializationError(t *testing.T) {
	assert := assert2.New(t)

	ranking := GetRankingInstance()
	defer ranking.Release()
	assert.NotNil(ranking)
	assert.False(ranking.isInitialized)

	players, err := ranking.GetPlayers()
	assert.Nil(players)

	if assert.NotNil(err) {
		assert.Equal(custom_errors.RankingInitializationErrorMessage, err.Error())
	}
}

// Unit Test: Checking adding players on top 10
// Partition Share: If i can update a player, i can update any
func TestInitializeRankingMultiplesPlayers(t *testing.T) {
	assert := assert2.New(t)

	ranking := GetRankingInstance()
	defer ranking.Release()
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

// Unit Test:
// Partition Share:
func TestGetRankingInstanceMultithreading(t *testing.T) {
	assert := assert2.New(t)
	var rankings [2]*Ranking
	var wg sync.WaitGroup

	for i := 0; i < len(rankings); i++ {
		wg.Add(1)
		go func(position int) {

			defer wg.Done()
			ranking := GetRankingInstance()
			assert.NotNil(ranking)
			rankings[position] = ranking
		}(i)
	}

	wg.Wait()
	assert.Equal(rankings[0], rankings[1])
}

// Unit Test: Checking adding players on top 10
// Partition Share: If i can add a player, i can add any
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
	ranking := GetRankingInstance()
	assert.NotNil(ranking)
	defer ranking.Release()

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

// Unit Test
func TestAddPlayersEmptyRanking(t *testing.T) {
	assert := assert2.New(t)

	for i := 0; i < maxPlayers; i++ {
		ranking := GetRankingInstance()
		assert.NotNil(ranking)

		if i == 0 { // To avoid collision with other test. But defer must be called once.
			defer ranking.Release()
		}

		assert.Equal(i, len(ranking.Players))

		newPlayer := Player{
			Name:   "test " + strconv.Itoa(i),
			Points: i + 1,
		}

		ranking.Update(newPlayer)
	}
}

// Unit Test
// Partition Share
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
	ranking := GetRankingInstance()
	assert.NotNil(ranking)

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
