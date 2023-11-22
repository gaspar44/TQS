package custom_errors

type UnknownDifficultyError struct {
}

func (err *UnknownDifficultyError) Error() string {
	return UnknownDifficultyErrorMessage
}
