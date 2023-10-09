package errors

type RankingInitializationError struct {
	internalError error
}

func NewRankingInitializationError(err error) *RankingInitializationError {
	return &RankingInitializationError{internalError: err}
}

func (err *RankingInitializationError) Error() string {
	return err.internalError.Error()
}
