package fortesting

import "testing"

// Test for add command
func TestData_AddFunc(t *testing.T) {

	var tests = []struct {
		input0 Data
		input1 string

		want string
	}{
		{Data{}, "add", "add added successfully"},
	}

	for _, test := range tests {
		got := test.input0.AddFunc(test.input1)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}

// Test for list command
func TestListData(t *testing.T) {
	var tests = []struct {
		want string
	}{
		{"list successfully printed"},
	}
	for _, test := range tests {
		got := ListData()
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}

// Test for done command
func TestData_DoneTask(t *testing.T) {
	var tests = []struct {
		input0 Data
		input1 string

		want string
	}{
		{Data{}, "apple", "Successful"},
	}
	for _, test := range tests {
		got := test.input0.DoneTask(test.input1)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}

//Test for undone command

func TestData_UndoneItem(t *testing.T) {
	var tests = []struct {
		input0 Data
		input1 string

		want string
	}{
		{Data{}, "apple", "Successful"},
	}
	for _, test := range tests {
		got := test.input0.UndoneItem(test.input1)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}

//Test for Cleanup command

func TestCleanup(t *testing.T) {
	var tests = []struct {
		want string
	}{
		{"Successful"},
	}
	for _, test := range tests {
		got := Cleanup()
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
