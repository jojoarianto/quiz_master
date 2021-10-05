package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/divan/num2words"
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/domain/repository"
	"github.com/jojoarianto/quiz_master/helper"
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

// Get all question to retrieve a data of all question
func (ps *questionService) GetAll() (question []model.Question, err error) {
	question, err = ps.questionRepo.GetAll()
	if err != nil {
		return
	}

	return
}

// Get question to retrieve a data of a question
func (ps *questionService) GetByNumber(number int) (question model.Question, err error) {
	question, err = ps.questionRepo.GetByNumber(number)
	if err != nil {
		return
	}

	return
}

// Add question service to save a data of question
func (ps *questionService) Add(question model.Question) error {
	questionExist, err := ps.questionRepo.GetByNumber(question.Number)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) == false {
			return err
		}
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

// Update question to update data question
func (ps *questionService) Update(number int, newQuestion model.Question) error {
	existingQuestion, err := ps.questionRepo.GetByNumber(number)
	if err != nil {
		return err
	}

	if existingQuestion.ID == 0 {
		return model.ErrQuestionNotFound
	}

	err = ps.questionRepo.Update(number, newQuestion)
	if err != nil {
		return err
	}

	return nil
}

// Answer question to answer of question will return correct or in correct
func (ps *questionService) Answer(number int, userAnswer string) (bool, error) {
	question, err := ps.questionRepo.GetByNumber(number)
	if err != nil {
		return false, err
	}

	// to lower case
	userAnswer = strings.ToLower(userAnswer)
	userAnswer = helper.RemoveQuotes(userAnswer)

	// simple comparison checking
	if userAnswer == strings.ToLower(question.Answer) {
		return true, nil
	}

	// state to check question.answer is int or not
	isQuestionAnswerInt := helper.IsEligibleConvertToInteger(question.Answer)

	// if integer then convert it to word in case user answer question with word
	if isQuestionAnswerInt == true {

		// convert to int
		answerInInt, _ := strconv.Atoi(question.Answer)
		answerInWord := num2words.Convert(answerInInt)

		// convert - to space (tweenty-six to tweenty six)
		answerInWord = strings.Replace(answerInWord, "-", " ", -1)

		// compare answer with question.answer in word
		if strings.ToLower(answerInWord) == userAnswer {
			return true, nil
		}

	}

	return false, nil
}
