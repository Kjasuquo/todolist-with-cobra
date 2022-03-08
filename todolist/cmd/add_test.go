package cmd

import "testing"

func TestGetData(t *testing.T) {
	var tests = []struct {
		input string

		want string
	}{
		{"add", "\n" + "add successfully added"},
		{"Pray today", "\n" + "Pray today successfully added"},
		{"Push task", "\n" + "Push task successfully added"},
	}

	for _, test := range tests {
		got := GetData(test.input)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
