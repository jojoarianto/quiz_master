package handler

import (
	"testing"

	"github.com/jojoarianto/quiz_master/domain/model"
)

func TestBuildViewQuestion(t *testing.T) {
	type args struct {
		question model.Question
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Scenario 1",
			args: args{model.Question{
				Number:   1,
				Question: "How many letter are in the English alphabet",
				Answer:   "26",
			}},
			want: "Q: How many letter are in the English alphabet\nA: 26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildViewQuestion(tt.args.question); got != tt.want {
				t.Errorf("BuildViewQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildViewAllQuestion(t *testing.T) {
	type args struct {
		questions []model.Question
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Scenario 1",
			args: args{
				[]model.Question{
					model.Question{
						Number:   1,
						Question: "How many letter are in the English alphabet",
						Answer:   "26",
					},
					model.Question{
						Number:   2,
						Question: "How old you are?",
						Answer:   "7",
					},
				},
			},
			want: "No | Question | Answer\n1. \"How many letter are in the English alphabet\" 26\n2. \"How old you are?\" 7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildViewAllQuestion(tt.args.questions); got != tt.want {
				t.Errorf("BuildViewAllQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
