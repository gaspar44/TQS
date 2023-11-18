package storage

import "gaspar44/TQS/model"

type StorageHandler interface {
	ReadRanking() (*model.Ranking, error)
	WriteRanking() error
}
