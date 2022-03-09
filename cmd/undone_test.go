package cmd

import "testing"

func TestUndone(t *testing.T) {
	var tests = []struct {
		input string

		want string
	}{
		{"Praying", "\n" + "Task has been undeleted"},
		{"Eating", "\n" + "Task has been undeleted"},
		{"Solving Algorithms", "\n" + "Task has been undeleted"},
	}
	for _, test := range tests {
		got := Undone(test.input)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
