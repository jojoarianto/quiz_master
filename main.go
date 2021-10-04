package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		err = CommandRouter(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func CommandRouter(cmdStr string) error {
	arrCommandStr := strings.Fields(cmdStr)

	switch arrCommandStr[0] {
	case "questions":
		handler.ShowAllQuestionHandler()
	case "question":
		handler.ShowQuestionHandler(cmdStr)
	case "answer_question":
		handler.AnswerQuestionHandler(cmdStr)
	case "create_question":
		handler.AddQuestionHandler(cmdStr)
	case "update_question":
		handler.UpdateQuestionHandler(cmdStr)
	case "delete_question":
		handler.DeleteQuestionHandler(cmdStr)
	case "help":
		fmt.Println(handler.BuildHelpMenu())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command not found, type \"help\" if you want to see available command")
	}

	fmt.Println()
	return nil
}

func AddCQue() {

}
