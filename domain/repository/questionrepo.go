package repository

import "github.com/jojoarianto/quiz_master/domain/model"

type QuestionRepo interface {
	Add(model.Question) error
	GetByNumber(number int) (question model.Question, x error)
}
