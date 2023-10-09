package storage

type StorageHandler interface {
	ReadRanking() ([]byte, error)
	WriteRanking() error
}
