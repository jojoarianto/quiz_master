package service

import (
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/domain/repository"
)

type questionService struct {
	questionRepo repository.QuestionRepo
}

// NewQuestionService method service init
func NewQuestionService(questionRepo repository.QuestionRepo) *questionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

// Add question service to save a data of question
func (ps *questionService) Add(question model.Question) error {
	questionExist, err := ps.questionRepo.GetByNumber(question.Number)
	if err != nil {
		return err
	}

	if questionExist.ID != 0 {
		return model.ErrQuestionAlreadyExist
	}

	err = ps.questionRepo.Add(question)
	if err != nil {
		return err
	}
	return nil
}

// Delete question service to delete a data of question
func (ps *questionService) Delete(number int) error {
	questionExist, err := ps.questionRepo.GetByNumber(number)
	if err != nil {
		return err
	}

	if questionExist.ID == 0 {
		return model.ErrQuestionNotFound
	}

	err = ps.questionRepo.Delete(number)
	if err != nil {
		return err
	}

	return nil
}
