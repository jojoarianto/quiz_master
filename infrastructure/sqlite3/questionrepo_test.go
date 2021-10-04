package sqlite3

import (
	"testing"

	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "../../test-db.sqlite3"
	Dialeg    = "sqlite3"
)

func Test_questionRepo_Update(t *testing.T) {
	// connect to test-db
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// reset test-db data
	db.DropTable(&model.Question{})
	db.AutoMigrate(&model.Question{})

	// test connection
	assert.NoError(t, err)

	// prepare data
	originalQuestion := model.Question{
		Number:   1,
		Question: "How many letter are in the English alphabet",
		Answer:   "26",
	}

	repo := NewQuestionRepo(db)
	err = repo.Add(originalQuestion)
	assert.NoError(t, err)

	t.Run("Fail Update Scenario", func(t *testing.T) {
		newQuestion := model.Question{
			Number:   100, // not exist number
			Question: "[updated] How many letter are in the English alphabet",
			Answer:   "27",
		}

		err = repo.Update(100, newQuestion)
		assert.NoError(t, err)

		// check data resultQuestion must equal with originalQuestion
		resultQuestion, err := repo.GetByNumber(1)
		assert.NoError(t, err)
		assert.Equal(t, originalQuestion.Question, resultQuestion.Question)
		assert.Equal(t, originalQuestion.Answer, resultQuestion.Answer)
	})

	t.Run("Success Update Scenario", func(t *testing.T) {
		newQuestion := model.Question{
			Number:   1,
			Question: "[updated] How many letter are in the English alphabet",
			Answer:   "27",
		}

		err = repo.Update(1, newQuestion)
		assert.NoError(t, err)

		// check number 1 must update as newQuestion object
		resultQuestion, err := repo.GetByNumber(1)
		assert.NoError(t, err)
		assert.Equal(t, newQuestion.Question, resultQuestion.Question)
		assert.Equal(t, newQuestion.Answer, resultQuestion.Answer)
	})

}
