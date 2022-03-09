package cmd

import "testing"

func TestCleanup(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"All done tasks have been deleted"},
	}
	for _, test := range tests {
		got := Cleanup()
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
