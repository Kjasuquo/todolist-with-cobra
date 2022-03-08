package main

import "testing"

func TestGetData(t *testing.T) {
	var tests = []struct {
		input string

		want string
	}{
		{"add", "add successfully added"},
	}

	for _, test := range tests {
		got := GetData(test.input)
		if got != test.want {
			t.Errorf("")
		}
	}
}

func TestRead(t *testing.T) {

}