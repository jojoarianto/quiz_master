package model

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")

	// ErrTrxFail will throw message that transaction fail
	ErrTrxFail = errors.New("Given Param is not valid")

	// ErrQuestionAlreadyExist will throw message question number already exist
	ErrQuestionAlreadyExist = errors.New("Question number already exist")
)
