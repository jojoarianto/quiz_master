package service

import "github.com/jojoarianto/quiz_master/domain/model"

// QuestionService contract
type QuestionService interface {
	Add(model.Question) error
	Delete(number int) error
}
