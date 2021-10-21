package arrays

import (
	"errors"
	"log"
	"reflect"
)

// ValidateArray - validates that data types are same
func ValidateArray(input1, input2 interface{}) (err error) {
	actualType := reflect.TypeOf(input1)
	expectedType := reflect.TypeOf(input2)

	if actualType.Kind() != expectedType.Kind() {
		log.Printf("Actual type [%s] is not same as expected type [%s]\n", actualType.Name(), expectedType.Name())
		err = errors.New(MismatchedTypeError)
	}

	return
}

// CheckExpectedType - validates that data types are same
func CheckExpectedType(validatedInput interface{}, expectedKind reflect.Kind) (err error) {
	inputType := reflect.TypeOf(validatedInput)

	if inputType.Kind() != expectedKind {
		log.Printf("Input type [%s] is not same as exptected type [%s]\n", inputType.Name(), expectedKind.String())
		err = errors.New(MismatchedTypeError)
	}

	return
}
