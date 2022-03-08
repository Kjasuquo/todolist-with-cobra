package cmd

import "testing"

func TestDone(t *testing.T) {
	tests := []struct {
		input string

		want string
	}{
		{"washing car", "\n" + "washing car has been successfully done"},
		{"Cooking", "\n" + "Cooking has been successfully done"},
		{"Ironing my clothes", "\n" + "Ironing my clothes has been successfully done"},
	}
	for _, test := range tests {
		got := Done(test.input)
		if got != test.want {
			t.Errorf("Execpected %v but got %v", test.want, got)
		}
	}
}
