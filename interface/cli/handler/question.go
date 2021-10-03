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
		Question: arrCommandStr[2],
		Number:   number,
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

	// if not error
	if err == nil {
		fmt.Printf("Question no %d created\n", number)
		return
	}

	// if error
	if err == model.ErrQuestionAlreadyExist {
		fmt.Println(model.ErrQuestionAlreadyExist)
	} else {
		fmt.Println("Fail to insert to database input to database")
	}

	return
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
		// if err == model.ErrQuestionNotFound {
		fmt.Println(err.Error())
		// } else {
		// 	fmt.Println("Fail to delete question")
		// }

		return
	}

	fmt.Printf("Question no %d was deleted!\n", number)
}
