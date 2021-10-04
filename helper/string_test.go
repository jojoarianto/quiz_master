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

func Test_removeQuotes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Scenario 1",
			args: args{str: "\"How many letter are in the English alphabet\""},
			want: "How many letter are in the English alphabet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveQuotes(tt.args.str); got != tt.want {
				t.Errorf("removeQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEligibleConvertToInteger(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: "76"},
			want: true,
		},
		{
			args: args{str: "Tujuh Enam"},
			want: false,
		},
		{
			args: args{str: "Name of person"},
			want: false,
		},
		{
			args: args{str: "123556"},
			want: true,
		},
		{
			args: args{str: "123H56"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEligibleConvertToInteger(tt.args.str); got != tt.want {
				t.Errorf("IsEligibleConvertToInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
