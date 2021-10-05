package handler

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/quiz_master/domain/handler"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/helper"
	"github.com/jojoarianto/quiz_master/infrastructure/sqlite3"
	"github.com/jojoarianto/quiz_master/service"
)

type connection struct {
	Conn *gorm.DB
}

// NewCliHandler method handler interface init
func NewCliHandler(conn *gorm.DB) handler.QuestionHandler {
	return &connection{Conn: conn}
}

func (c *connection) ShowAllQuestionHandler() {
	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	questions, err := questionSvc.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := BuildViewAllQuestion(questions)
	fmt.Println(result)
}

func (c *connection) ShowQuestionHandler(cmdStr string) {
	arrCommandStr := helper.StringSplitter(cmdStr)
	if len(arrCommandStr) < 2 {
		fmt.Println("invalid input")
		return
	}

	number, err := strconv.Atoi(arrCommandStr[1])
	if err != nil {
		fmt.Println("number must be integer")
		return
	}

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	question, err := questionSvc.GetByNumber(number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := BuildViewQuestion(question)
	fmt.Println(result)
}

func (c *connection) AddQuestionHandler(cmdStr string) {
	arrCommandStr := helper.StringSplitter(cmdStr)
	if len(arrCommandStr) < 4 {
		fmt.Println("invalid input")
		return
	}

	number, err := strconv.Atoi(arrCommandStr[1])
	if err != nil {
		fmt.Println("number must be integer")
		return
	}

	question := model.Question{
		Question: helper.RemoveQuotes(arrCommandStr[2]),
		Number:   number,
		Answer:   helper.RemoveQuotes(arrCommandStr[3]),
	}

	// Todo: validate input

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	err = questionSvc.Add(question)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d created\n", number)
}

func (c *connection) DeleteQuestionHandler(cmdStr string) {
	arrCommandStr := helper.StringSplitter(cmdStr)
	if len(arrCommandStr) != 2 {
		fmt.Println("invalid input")
		return
	}

	number, err := strconv.Atoi(arrCommandStr[1])
	if err != nil {
		fmt.Println("number must be integer")
		return
	}

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	err = questionSvc.Delete(number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d was deleted!\n", number)
}

func (c *connection) UpdateQuestionHandler(cmdStr string) {
	arrCommandStr := helper.StringSplitter(cmdStr)
	if len(arrCommandStr) != 4 {
		fmt.Println("invalid input")
		return
	}

	number, err := strconv.Atoi(arrCommandStr[1])
	if err != nil {
		fmt.Println("number must be integer")
		return
	}

	question := model.Question{
		Question: helper.RemoveQuotes(arrCommandStr[2]),
		Number:   number,
		Answer:   helper.RemoveQuotes(arrCommandStr[3]),
	}

	// Todo: validate input

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	err = questionSvc.Update(number, question)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d updated\n", number)
}

func (c *connection) AnswerQuestionHandler(cmdStr string) {
	arrCommandStr := helper.StringSplitter(cmdStr)
	if len(arrCommandStr) != 3 {
		fmt.Println("invalid input")
		return
	}

	number, err := strconv.Atoi(arrCommandStr[1])
	if err != nil {
		fmt.Println("number must be integer")
		return
	}

	userAnswer := helper.RemoveQuotes(arrCommandStr[2])

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(c.Conn))
	isCorrect, err := questionSvc.Answer(number, userAnswer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if isCorrect == false {
		fmt.Println("Uncorrect")
		return
	}

	fmt.Println("Correct")
}
