package handler

import (
	"fmt"
	"strconv"

	"github.com/jojoarianto/quiz_master/domain/model"
)

func BuildViewQuestion(question model.Question) string {
	str := fmt.Sprintf("Q: %s\nA: %s", question.Question, question.Answer)
	return str
}

func BuildViewAllQuestion(questions []model.Question) string {
	str := "No | Question | Answer"
	for _, question := range questions {
		str = str + "\n" + strconv.Itoa(question.Number) + ". \"" + question.Question + "\" " + question.Answer
	}
	return str
}

func BuildHelpMenu() string {
	str := "Command | Description\n" +
		"create_question <no> <question> <answer> | Creates a question\n" +
		"update_question <no> <question> <answer> | Updates a question\n" +
		"delete_question <no> | Deletes a question\n" +
		"question <no> | Shows a question\n" +
		"questions | Shows question list"

	return str
}
