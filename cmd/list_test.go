package cmd

import "testing"

func TestRead(t *testing.T) {
	var tests = []struct {
		want string
	}{
		{"\n" + "list successfully printed"},
	}

	for _, test := range tests {
		got := Read()
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
