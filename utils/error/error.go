package error

import "gobrax-api/domain/entities"

// createError - returns error struct
func CreateError(errorMessage string) entities.Error {
	e := entities.Error{}
	e.ErrorMessage = errorMessage
	return e
}
