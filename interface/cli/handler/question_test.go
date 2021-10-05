package handler

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "test-db.sqlite3"
	Dialeg    = "sqlite3"
)

func initDatabaseTest() *gorm.DB {

	// init database for testing
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		panic(err)
	}

	// reset test-db data
	db.AutoMigrate(&model.Question{})
	db.DropTable(&model.Question{})
	db.AutoMigrate(&model.Question{})

	return db
}

// catureOutput to capture stdOut from console
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer

	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()
	f()
	writer.Close()

	return <-out
}

func TestAddQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// scenario 1
	output := captureOutput(func() {
		cliHandler.AddQuestionHandler("create_question 1 \"How many alphabet in english\"")
	})
	assert.Equal(t, "invalid input\n", output)

	// scenario 2
	output = captureOutput(func() {
		cliHandler.AddQuestionHandler("create_question abc \"How many alphabet in english\" 26")
	})
	assert.Equal(t, "number must be integer\n", output)

	// scenario 3
	output = captureOutput(func() {
		cliHandler.AddQuestionHandler("create_question 4 \"How many alphabet in english\" 26")
	})
	assert.Equal(t, "Question no 4 created\n", output)

	// scenario 4
	output = captureOutput(func() {
		cliHandler.AddQuestionHandler("create_question 4 \"How many alphabet in english\" 26")
	})
	assert.Equal(t, "Question number already exist\n", output)
}

func Test_connection_ShowAllQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// without data
	output := captureOutput(func() {
		cliHandler.ShowAllQuestionHandler()
	})
	assert.Equal(t, "No | Question | Answer\n", output)

	// prepare data
	cliHandler.AddQuestionHandler("create_question 1 \"How many letter are in the English alphabet\" 26")
	cliHandler.AddQuestionHandler("create_question 2 \"President Name\" Jokowi")

	// with data
	output = captureOutput(func() {
		cliHandler.ShowAllQuestionHandler()
	})
	assert.Equal(t, "No | Question | Answer\n1. \"How many letter are in the English alphabet\" 26\n2. \"President Name\" Jokowi\n", output)
}

func Test_connection_AnswerQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// without data
	output := captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question 6 26")
	})
	assert.Equal(t, "record not found\n", output)

	// prepare data
	cliHandler.AddQuestionHandler("create_question 99 \"How many letter are in the English alphabet\" 26")

	// with data
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question 99 26")
	})
	assert.Equal(t, "Correct\n", output)

	// with data and answer with word
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question 99 \"twenty six\"")
	})
	assert.Equal(t, "Correct\n", output)

	// wrong answer
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question 99 9")
	})
	assert.Equal(t, "Uncorrect\n", output)

	// invalid input
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question")
	})
	assert.Equal(t, "invalid input\n", output)

	// invalid type of number
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question sembilan 9")
	})
	assert.Equal(t, "number must be integer\n", output)

	// not exist number
	output = captureOutput(func() {
		cliHandler.AnswerQuestionHandler("answer_question 1000 9")
	})
	assert.Equal(t, "record not found\n", output)
}

func Test_connection_DeleteQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// to much argument
	output := captureOutput(func() {
		cliHandler.DeleteQuestionHandler("delete_question 1 \"How many alphabet in english\"")
	})
	assert.Equal(t, "invalid input\n", output)

	// input must be integer
	output = captureOutput(func() {
		cliHandler.DeleteQuestionHandler("delete_question hello")
	})
	assert.Equal(t, "number must be integer\n", output)

	// without data
	output = captureOutput(func() {
		cliHandler.DeleteQuestionHandler("delete_question 7")
	})
	assert.Equal(t, "record not found\n", output)

	// prepare data
	cliHandler.AddQuestionHandler("create_question 77 \"How many letter are in the English alphabet\" 26")

	// with data
	output = captureOutput(func() {
		cliHandler.DeleteQuestionHandler("delete_question 77")
	})
	assert.Equal(t, "Question no 77 was deleted!\n", output)
}

func Test_connection_UpdateQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// prepare data
	cliHandler.AddQuestionHandler("create_question 1 \"How many alphabet in english\" 26")

	// to much argument
	output := captureOutput(func() {
		cliHandler.UpdateQuestionHandler("update_question 1")
	})
	assert.Equal(t, "invalid input\n", output)

	// input must be integer
	output = captureOutput(func() {
		cliHandler.UpdateQuestionHandler("update_question abc \"How many alphabet in english\" 77")
	})
	assert.Equal(t, "number must be integer\n", output)

	// with data
	output = captureOutput(func() {
		cliHandler.UpdateQuestionHandler("update_question 1 \"How many alphabet in english\" 77")
	})
	assert.Equal(t, "Question no 1 updated\n", output)

}

func Test_connection_ShowQuestionHandler(t *testing.T) {

	// create cliHandler
	cliHandler := NewCliHandler(initDatabaseTest())

	// prepare data
	cliHandler.AddQuestionHandler("create_question 1 \"How many alphabet in english\" 26")

	// with data
	output := captureOutput(func() {
		cliHandler.ShowQuestionHandler("question 1")
	})
	assert.Equal(t, "Q: How many alphabet in english\nA: 26\n", output)

	// must integer
	output = captureOutput(func() {
		cliHandler.ShowQuestionHandler("question abc")
	})
	assert.Equal(t, "number must be integer\n", output)

	// invalid input
	output = captureOutput(func() {
		cliHandler.ShowQuestionHandler("question")
	})
	assert.Equal(t, "invalid input\n", output)

}
