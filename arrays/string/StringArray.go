package string

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays"
	"log"
	"reflect"
)

type StringArray []string

func (r StringArray) IsPresent(searchedValue interface{}) (bool, error) {
	if reflect.TypeOf(searchedValue).Kind() != reflect.String {
		log.Println("Parameter type is not of string type")
		return false, errors.New(arrays.MismatchedTypeError)
	}

	for index := range r {
		if r[index] == searchedValue {
			return true, nil
		}
	}

	return false, nil
}

func (r StringArray) FindIndex(searchedValue interface{}) (int, error) {
	if reflect.TypeOf(searchedValue).Kind() != reflect.String {
		log.Println("Parameter type is not of string type")
		return 0, errors.New(arrays.MismatchedTypeError)
	}

	for index := range r {
		if r[index] == searchedValue {
			return index, nil
		}
	}

	return 0, nil
}

func (r StringArray) Get(index interface{}) (interface{}, error) {
	indexValue, ok := index.(int)
	if !ok {
		log.Println("Parameter type is not of Int type")
		return 0, errors.New(arrays.MismatchedTypeError)
	}

	if indexValue > len(r) || indexValue < 0 {
		log.Println("Provided index parameter is out of bounds")
		return 0, errors.New("index out of bounds")
	} else {
		return r[indexValue], nil
	}
}
