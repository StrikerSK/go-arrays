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

func (r StringArray) Get(searchedValue interface{}) (interface{}, error) {
	index, err := r.FindIndex(searchedValue)
	if err != nil {
		return nil, err
	}

	if index == 0 {
		return nil, errors.New("could not find searched structure")
	} else {
		return r[index], nil
	}
}
