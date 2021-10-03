package helper

import (
	"reflect"
	"testing"
)

func TestStringSplitter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Scenario 1",
			args: args{str: "create_question 1 \"How many letters are there in the English alphabet?\" 26"},
			want: []string{"create_question", "1", "\"How many letters are there in the English alphabet?\"", "26"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSplitter(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSplitter() = %v, want %v", got, tt.want)
			}
		})
	}
}
