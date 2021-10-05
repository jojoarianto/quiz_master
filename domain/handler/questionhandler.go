package handler

// QuestionService contract
type QuestionHandler interface {
	ShowAllQuestionHandler()
	ShowQuestionHandler(cmdStr string)
	AddQuestionHandler(cmdStr string)
	AnswerQuestionHandler(cmdStr string)
	DeleteQuestionHandler(cmdStr string)
	UpdateQuestionHandler(cmdStr string)
}
