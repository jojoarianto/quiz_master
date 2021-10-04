package service

import (
	"testing"

	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/domain/model"
	"github.com/jojoarianto/quiz_master/domain/repository"
	"github.com/jojoarianto/quiz_master/infrastructure/sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "../test-db.sqlite3"
	Dialeg    = "sqlite3"
)

func Test_questionService_Answer(t *testing.T) {

	// connect to test-db
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// reset test-db data
	db.AutoMigrate(&model.Question{})
	db.DropTable(&model.Question{})
	db.AutoMigrate(&model.Question{})

	// test connection
	assert.NoError(t, err)

	// prepare data
	repo := sqlite3.NewQuestionRepo(db)

	// seeding data
	seedingData(repo)

	type fields struct {
		questionRepo repository.QuestionRepo
	}
	type args struct {
		number     int
		userAnswer string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 1, userAnswer: "26"},
			want:    true,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 1, userAnswer: "twenty six"},
			want:    true,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 1, userAnswer: "Twenty Six"},
			want:    true,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 1, userAnswer: "Tweunty S1x"},
			want:    false,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 1, userAnswer: "88"},
			want:    false,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 200, userAnswer: "10"},
			want:    false,
			wantErr: true, // question not found
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 2, userAnswer: "Bapak Jokowi"},
			want:    true,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 2, userAnswer: "Bapak jokowow"},
			want:    false,
			wantErr: false,
		},
		{
			fields:  fields{questionRepo: repo},
			args:    args{number: 2, userAnswer: "76"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &questionService{
				questionRepo: tt.fields.questionRepo,
			}
			got, err := ps.Answer(tt.args.number, tt.args.userAnswer)
			if (err != nil) != tt.wantErr {
				t.Errorf("questionService.Answer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("questionService.Answer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func seedingData(repo repository.QuestionRepo) {
	originalQuestion := model.Question{
		Number:   1,
		Question: "How many letter are in the English alphabet",
		Answer:   "26",
	}
	repo.Add(originalQuestion)

	originalQuestion = model.Question{
		Number:   2,
		Question: "Current Presiden Name",
		Answer:   "Bapak Jokowi",
	}
	repo.Add(originalQuestion)
}
