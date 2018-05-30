package words

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"basic",
			args{"-- .! I am? what I am!!"},
			[]string{"i", "am", "what", "i", "am"},
		},
		{
			"splits with suffixes",
			args{"I'm Here"},
			[]string{"i'm", "here"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Words(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Words() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveWords(t *testing.T) {
	type args struct {
		s     string
		words []string
		repl  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"removes words",
			args{"Hello my dear friend", []string{"my"}, "!"},
			"Hello ! dear friend",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveWords(tt.args.s, tt.args.words, tt.args.repl); got != tt.want {
				t.Errorf("RemoveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
