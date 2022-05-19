package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	input = strings.TrimSpace(input)

	// Check if input emty
	if len(input) == 0 {
		return "", fmt.Errorf("[StringSum] level 1 error: %w", errorEmptyInput)
	}
	var (
		subStringsArray     []string
		numberOne           int
		numberTwo           int
		ifNumberOneNegative int
	)
	countPlus := strings.Count(input, "+")

	// if string include two number and "+"
	if countPlus > 1 {
		return "", fmt.Errorf("[StringSum] level 1 error: %w", errorNotTwoOperands)
	}
	if countPlus == 1 {
		subStringsArray = strings.Split(input, "+")
		if len(subStringsArray) != 2 {
			return "", fmt.Errorf("[StringSum] level 1 error: %w", errorNotTwoOperands)
		}
		if len(subStringsArray) == 2 {
			numberOne, err = strconv.Atoi(strings.TrimSpace(subStringsArray[0]))
			if err != nil {
				return "", fmt.Errorf("[StringSum] level 1 error: %w", err.(*strconv.NumError))
			}
			numberTwo, _ = strconv.Atoi(strings.TrimSpace(subStringsArray[1]))
			if err != nil {
				return "", fmt.Errorf("[StringSum] level 1 error: %w", err.(*strconv.NumError))
			}

			return strconv.Itoa(numberOne + numberTwo), nil
		}

	}
	// chekFirstelement
	if input[0] == '-' {
		input = input[1:]
		ifNumberOneNegative = -1
	} else {
		ifNumberOneNegative = 1
	}
	countMinus := strings.Count(input, "-")

	if countMinus == 1 || countMinus == 2 {
		subStringsArray = strings.Split(input, "-")
		if len(subStringsArray) == 2 {
			numberOne, err = strconv.Atoi(strings.TrimSpace(subStringsArray[0]))
			if err != nil {
				return "", fmt.Errorf("[StringSum] level 1 error: %w", err.(*strconv.NumError))
			}
			numberTwo, err = strconv.Atoi(strings.TrimSpace(subStringsArray[1]))
			if err != nil {
				return "", fmt.Errorf("[StringSum] level 1 error: %w", err.(*strconv.NumError))
			}

		}
		return strconv.Itoa(ifNumberOneNegative*numberOne - numberTwo), nil
	}
	return "", fmt.Errorf("[StringSum] level 1 error: %w", err.(*strconv.NumError))
}
