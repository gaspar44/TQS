package custom_errors

import "strconv"

type InvalidCardPositionError struct {
	message string
}

func NewInvalidPositionError(position int) *InvalidCardPositionError {
	stringPosition := strconv.Itoa(position)
	return &InvalidCardPositionError{
		message: GameAlreadyInitializedErrorMessage + stringPosition,
	}
}
func (err *InvalidCardPositionError) Error() string {
	return err.message
}
