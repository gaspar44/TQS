package custom_errors

type GameAlreadyInitializedError struct {
}

func (err *GameAlreadyInitializedError) Error() string {
	return GameAlreadyInitializedErrorMessage
}
