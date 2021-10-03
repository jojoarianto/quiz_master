package handler

import (
	"fmt"

	"github.com/jojoarianto/quiz_master/domain/model"
)

func PrintQuestion(model.Question) {
	fmt.Println("Q: ")
	fmt.Println("A: ")
}
