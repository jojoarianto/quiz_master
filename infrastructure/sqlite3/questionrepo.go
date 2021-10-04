package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/domain/repository"
)

type questionRepo struct {
	Conn *gorm.DB
}

// NewQuestionRepo method repo init
func NewQuestionRepo(conn *gorm.DB) repository.QuestionRepo {
	return &questionRepo{Conn: conn}
}

// Add method to add new question
func (qr *questionRepo) Add(question model.Question) error {
	if err := qr.Conn.Save(&question).Error; err != nil {
		return err
	}

	return nil
}

// GetByNumber to get single question by number
func (qr *questionRepo) GetByNumber(number int) (question model.Question, err error) {
	err = qr.Conn.Where("number = ?", number).First(&question).Error
	if err != nil {
		return
	}

	return
}

// Add method to get all question
func (qr *questionRepo) GetAll() (questions []model.Question, err error) {
	if err = qr.Conn.Order("number asc").Find(&questions).Error; err != nil {
		return
	}
	return
}

// Delete method to delete a question
func (qr *questionRepo) Delete(number int) error {
	if err := qr.Conn.Where("number = ?", number).Delete(model.Question{}).Error; err != nil {
		return err
	}

	return nil
}
