package handler

import (
	"fmt"

	"github.com/jojoarianto/quiz_master/domain/model"
)

func PrintQuestion(question model.Question) {
	fmt.Printf("Q: %s", question.Question)
	fmt.Println()
	fmt.Printf("A: %s", question.Answer)
	fmt.Println()
}
