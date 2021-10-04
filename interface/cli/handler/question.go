package handler

import (
	"fmt"
	"strconv"

	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/helper"
	"github.com/jojoarianto/quiz_master/infrastructure/sqlite3"
	"github.com/jojoarianto/quiz_master/service"
)

func ShowAllQuestionHandler() {
	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
	questions, err := questionSvc.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := BuildViewAllQuestion(questions)
	fmt.Println(result)
}

func ShowQuestionHandler(cmdStr string) {
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

	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
	question, err := questionSvc.GetByNumber(number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := BuildViewQuestion(question)
	fmt.Println(result)
}

func AddQuestionHandler(cmdStr string) {
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

	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
	err = questionSvc.Add(question)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d created\n", number)
}

func DeleteQuestionHandler(cmdStr string) {
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

	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
	err = questionSvc.Delete(number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d was deleted!\n", number)
}

func UpdateQuestionHandler(cmdStr string) {
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

	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	question := model.Question{
		Question: helper.RemoveQuotes(arrCommandStr[2]),
		Number:   number,
		Answer:   helper.RemoveQuotes(arrCommandStr[3]),
	}

	// Todo: validate input

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
	err = questionSvc.Update(number, question)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Question no %d updated\n", number)
}

func AnswerQuestionHandler(cmdStr string) {
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

	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println(model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	questionSvc := service.NewQuestionService(sqlite3.NewQuestionRepo(db))
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
