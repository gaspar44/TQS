package custom_errors

import "errors"

type RankingInitializationError struct {
	internalError error
}

func NewRankingInitializationError(err error) *RankingInitializationError {
	return &RankingInitializationError{internalError: err}
}

func NewRankingInitializationErrorWithMessage() *RankingInitializationError {
	return &RankingInitializationError{internalError: errors.New(RankingInitializationErrorMessage)}
}

func (err *RankingInitializationError) Error() string {
	return err.internalError.Error()
}
