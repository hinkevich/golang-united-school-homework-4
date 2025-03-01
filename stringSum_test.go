package string_sum

import (
	"fmt"
	"testing"
)

type testData struct {
	inputString string
	want        string
	wantError   error
}

func TestStringSummValidInput(t *testing.T) {
	tests := []testData{
		{inputString: "3+5", want: "8", wantError: nil},
		{inputString: "-3+5", want: "2", wantError: nil},
		{inputString: "3 + 5", want: "8", wantError: nil},
		{inputString: "-3 + -5", want: "-8", wantError: nil},
		{inputString: " -3 + -5", want: "-8", wantError: nil},
	}
	for _, test := range tests {
		got, _ := StringSum(test.inputString)
		if got != test.want {
			t.Errorf("test summ valid Input - failed")
		}

	}
}
func TestStringDifferenceValidInput(t *testing.T) {
	tests := []testData{

		{inputString: "3-5", want: "-2", wantError: nil},
		{inputString: "-3-5", want: "-8", wantError: nil},
		{inputString: "-3 -5", want: "-8", wantError: nil},

		{inputString: " -3 -5 ", want: "-8", wantError: nil},
	}
	for _, test := range tests {
		got, _ := StringSum(test.inputString)
		if got != test.want {
			t.Errorf("test difference valid Input - failed")
		}

	}
}
func TestStringInvalidInput(t *testing.T) {
	tests := []testData{
		{inputString: "3+5+8", want: "", wantError: fmt.Errorf("[StringSum] level 1 error: %w", errorNotTwoOperands)},
		{inputString: "-3+5+6", want: "", wantError: fmt.Errorf("[StringSum] level 1 error: %w", errorNotTwoOperands)},
		{inputString: "", want: "", wantError: fmt.Errorf("[StringSum] level 1 error: %w", errorEmptyInput)},
	}
	for _, test := range tests {
		got, err := StringSum(test.inputString)
		if got != test.want && err != test.wantError {
			t.Errorf("sdfsdf")
		}

	}
}
