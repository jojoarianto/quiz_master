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

// Add method to add new question
func (qr *questionRepo) GetByNumber(number int) (question model.Question, x error) {
	err := qr.Conn.Where("number = ?", number).First(&question)
	if err != nil {
		return
	}

	return
}
