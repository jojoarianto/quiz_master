package handler

import (
	"fmt"

	"github.com/jojoarianto/quiz_master/domain/model"
)

func BuildViewQuestion(question model.Question) string {
	str := fmt.Sprintf("Q: %s\nA: %s", question.Question, question.Answer)
	return str
}
