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
		Answer:   arrCommandStr[3],
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
