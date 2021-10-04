package service

import "github.com/jojoarianto/quiz_master/domain/model"

// QuestionService contract
type QuestionService interface {
	Add(model.Question) error
	Delete(number int) error
	Update(number int, question model.Question) error
	GetAll() (question []model.Question, err error)
	GetByNumber(number int) (question model.Question, err error)
	Answer(number int, answer string) (isCorrect bool, err error)
}
