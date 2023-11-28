package model

import (
	"gaspar44/TQS/model/custom_errors"
	"strconv"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

// Unit test: Checking "NewGame()" function (easy)
// Test de cobertura (easy)
func TestNewGameEasyMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < EasyDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

func TestNewGameUnknownDifficulty(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Difficulty(3))
	assert.Nil(game)
	assert.NotNil(err)
	assert.Equal(custom_errors.UnknownDifficultyErrorMessage, err.Error())
}

// Unit test: Checking "NewGame()" function (Medium)
// Test de cobertura (easy)
func TestNewGameMediumMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Medium)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(MediumDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.GetPoints())
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < MediumDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

// Unit test: Checking "NewGame()" function (Hard)
// Test de cobertura (easy)
func TestNewGameHardMode(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(HardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.True(game.initialized)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	expectedElements := make(map[int]int)

	for i := 0; i < HardDifficultyCardsTotal/2; i++ {
		expectedElements[i] = 2
	}

	countedElements := countDuplicates(game.GetCards())
	assert.Equal(expectedElements, countedElements)
}

// Unit test: Checking "shuffleCards()" functions (Easy)
// Loop Test ->
func TestGameEasyModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Easy)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

// Unit test: Checking "shuffleCards()" functions (Medium)
// Loop Test ->
func TestGameMediumModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Medium)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(MediumDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

// Unit test: Checking "shuffleCards()" functions (Hard)
// Loop Test ->
func TestGameHardModeShuffle(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test2"
	game, err := NewGame(playerName, Hard)

	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(HardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	actualCards := make([]Card, len(game.GetCards()))
	copy(actualCards, game.GetCards())
	game.shuffleCards()
	shuffledCards := game.GetCards()

	assert.Equal(len(actualCards), len(shuffledCards))
	assert.True(checkCardsOrder(actualCards, shuffledCards))
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card  (easy)
// Frontier
func TestGameEasyModeSingleChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
}

// Frontier
func TestGameEasyModeFrontierUpperChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(EasyDifficultyCardsTotal - 1)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 2 correct cards (easy)
func TestGameEasyModeSameCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.Equal(-1, game.selectedCard.Position)

	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Equal(newSelectedCard, game.selectedCard)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card  (Hard)
func TestGameHardModeSingleChooseCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.True(game.initialized)
	assert.Equal(HardDifficultyCardsTotal, len(game.GetCards()))
	assert.Equal(0, game.points)
	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(4)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard

	assert.NotEqual(newSelectedCard, previousSelectedCard)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (easy)
func TestGameEasyModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(4)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(easyDifficultyPenalization+1, game.points)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (Medium)
func TestGameMediumModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Medium)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(MediumDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Medium)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(5)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(mediumDifficultyPenalization+1, game.points)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (Hard)
func TestGameHardModeWrongCardChose(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(HardDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Hard)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(5)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)
	assert.Equal(hardDifficultyPenalization+1, game.points)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (easy)
// Decision coverage
func TestGameHardModeChoseSameCardTwice(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Hard)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(HardDifficultyCardsTotal, len(game.GetCards()))

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	sameCardChoice := game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.Equal(sameCardChoice, newSelectedCard)
	assert.Equal(1, game.points)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (easy)
func TestGameCorrectCards(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(1)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	assert.Equal(initializationCard.Card.GetValue(), newSelectedCard.Card.GetValue())
	assert.Equal(initializationCard.Position, newSelectedCard.Position)
	assert.True(previousSelectedCard.Card.IsDisable)
	assert.True(game.GetCards()[1].IsDisable)
	assert.Equal(1, game.points)
}

// Unit test: Checking "ChooseCardOnBoard()" function for 1 card (easy)
func TestGameSelectDisabledCard(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))

	hackedCardList := createCards(Easy)
	game.cards = hackedCardList // Only to obtain a card list. Only because it's go's testing, we can access and modify private attributes

	assert.Equal(-1, game.selectedCard.Card.Value)
	assert.False(game.selectedCard.Card.IsVisible)
	assert.False(game.selectedCard.Card.IsDisable)
	assert.Equal(-1, game.selectedCard.Position)

	previousSelectedCard := game.selectedCard
	correctCards, err := game.ChooseCardOnBoard(0)
	assert.Nil(err)
	assert.True(correctCards)
	newSelectedCard := game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	previousSelectedCard = game.selectedCard
	correctCards, err = game.ChooseCardOnBoard(1)
	assert.True(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard
	assert.NotEqual(newSelectedCard, previousSelectedCard)

	correctCards, err = game.ChooseCardOnBoard(1)
	assert.False(correctCards)
	assert.Nil(err)
	newSelectedCard = game.selectedCard

	assert.Equal(-1, newSelectedCard.Card.Value)
	assert.False(newSelectedCard.Card.IsVisible)
	assert.False(newSelectedCard.Card.IsDisable)
	assert.Equal(-1, newSelectedCard.Position)
}

// Unit Test: Checking "Stop()" function
func TestGameStop(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)

	ranking := GetRankingInstance()
	assert.NotNil(ranking)
	defer ranking.Release()
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
	game.Stop()

	setPlayers, err := ranking.GetPlayers()
	assert.Nil(err)
	assert.Equal(players, setPlayers)
}

// Unit Test: Checking card position
// Limit values lower
// Decision
func TestGameInvalidLowerCardSelection(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))

	invalidPosition := -1
	correctCards, err := game.ChooseCardOnBoard(invalidPosition)
	assert.False(correctCards)
	assert.NotNil(err)
	assert.Equal(err.Error(), custom_errors.InvalidCardPositionErrorMessage+strconv.Itoa(invalidPosition))
}

// Unit Test: Checking card position
// Limit values upper
// Decision
func TestGameInvalidUpperCardSelection(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))

	invalidPosition := 45

	correctCards, err := game.ChooseCardOnBoard(invalidPosition)
	assert.False(correctCards)
	assert.NotNil(err)
	assert.Equal(err.Error(), custom_errors.InvalidCardPositionErrorMessage+strconv.Itoa(invalidPosition))
}

func TestGameAlreadyInitializedError(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	assert.Equal(EasyDifficultyCardsTotal, len(game.GetCards()))
	assert.True(game.initialized)

	err = game.createCards(Easy)
	assert.NotNil(err)
	assert.Equal(err.Error(), custom_errors.GameAlreadyInitializedErrorMessage)
}

func TestShuffleWithoutGameInitializationError(t *testing.T) {
	assert := assert2.New(t)
	playerName := "test1"

	game, err := NewGame(playerName, Easy)
	assert.Nil(err)
	assert.Equal(playerName, game.playerName)
	assert.Equal(0, game.points)
	game.initialized = false

	game.shuffleCards()
	assert.False(game.initialized)
}

func createCards(difficulty Difficulty) []Card {
	cards := make([]Card, 0)

	switch difficulty {
	case Easy:
		for i := 0; i < EasyDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Medium:
		for i := 0; i < MediumDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	case Hard:
		for i := 0; i < HardDifficultyCardsTotal/2; i++ {
			newCard := NewCard(i)
			newCardPair := NewCard(i)
			cards = append(cards, newCard)
			cards = append(cards, newCardPair)
		}
	}

	return cards
}

func checkCardsOrder(beforeShuffleCards []Card, afterShuffleCards []Card) bool {
	for i := range beforeShuffleCards {
		if beforeShuffleCards[i].GetValue() != afterShuffleCards[i].GetValue() {
			return true
		}
	}
	return false
}

func countDuplicates(cards []Card) map[int]int {
	countedElements := make(map[int]int)

	for i := 0; i < len(cards); i++ {
		cardValue := cards[i].Value
		_, existsKey := countedElements[cardValue]

		if !existsKey {
			countedElements[cardValue] = 1
		} else {
			countedElements[cardValue] += 1
		}
	}
	return countedElements
}
