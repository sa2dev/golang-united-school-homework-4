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
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("an error occurred: 	%w", errorEmptyInput)
	}

	if strings.Contains(input, "+") {
		numbers := strings.Split(input, "+")
		return sumTwoInt(numbers)
	}

	if strings.HasPrefix(strings.TrimSpace(input), "-") {
		newInput := strings.Replace(input, "-", "", 1)
		println(newInput)
		newInput = strings.Replace(newInput, "-", "+", 1)
		println(newInput)
		numbers := strings.Split(newInput, "+")
		println(numbers)
		res, err := sumTwoInt(numbers)
		return "-" + res, err
	}

	return "", fmt.Errorf("an error occurred")
}

func sumTwoInt(input []string) (output string, err error) {

	if len(input) != 2 {
		return "", fmt.Errorf("an error occurred: 	%w", errorNotTwoOperands)
	}

	firstNumber, err := strconv.Atoi(strings.TrimSpace(input[0]))
	if err != nil {
		return "", fmt.Errorf("an error occurred: 	%w", err)
	}

	secondNumber, err := strconv.Atoi(strings.TrimSpace(input[1]))
	if err != nil {
		return "", fmt.Errorf("an error occurred: 	%w", err)
	}

	return strconv.Itoa(firstNumber + secondNumber), nil
}
