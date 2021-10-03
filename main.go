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
	case "create_question":
		handler.AddQuestionHandler(cmdStr)
	case "delete_question":
		handler.DeleteQuestionHandler(cmdStr)
	case "help":
		fmt.Println("Command | Description")
		fmt.Println("create_question <no> <question> <answer> | Creates a question")
		fmt.Println("update_question <no> <question> <answer> | Updates a question")
		fmt.Println("delete_question <no> | Deletes a question")
		fmt.Println("question <no> | Shows a question")
		fmt.Println("questions | Shows question list")
	case "exit":
		os.Exit(0)
	}

	return nil
}

func AddCQue() {

}
