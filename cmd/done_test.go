package cmd

import "testing"

func TestDone(t *testing.T) {
	tests := []struct {
		input string

		want string
	}{
		{"washing car", "Successfully done tasks"},
	}
	for _, test := range tests {
		got := Done(test.input)
		if got != test.want {
			t.Errorf("Execpected %v but got %v", test.want, got)
		}
	}
}
