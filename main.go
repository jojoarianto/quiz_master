package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/interface/cli/handler"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Quiz Master!")

	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if len(cmdString) <= 1 {
			// for handling empty command
			continue
		}
		err = Run(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Run(cmdStr string) error {

	// init database
	conf := config.NewConfig(config.Dialeg, config.URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// routing command
	CommandRouter(db, cmdStr)

	return nil
}

func CommandRouter(db *gorm.DB, cmdStr string) {
	cliHandler := handler.NewCliHandler(db)
	arrCommandStr := strings.Fields(cmdStr)

	switch arrCommandStr[0] {
	case "questions":
		cliHandler.ShowAllQuestionHandler()
	case "question":
		cliHandler.ShowQuestionHandler(cmdStr)
	case "answer_question":
		cliHandler.AnswerQuestionHandler(cmdStr)
	case "create_question":
		cliHandler.AddQuestionHandler(cmdStr)
	case "update_question":
		cliHandler.UpdateQuestionHandler(cmdStr)
	case "delete_question":
		cliHandler.DeleteQuestionHandler(cmdStr)
	case "help":
		fmt.Println(handler.BuildHelpMenu())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command not found, type \"help\" if you want to see available command")
	}

	fmt.Println()
}
