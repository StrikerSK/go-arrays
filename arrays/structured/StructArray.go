package structured

import (
	"errors"
	"log"
)

// StructArray - Array that can iterate structured types by defined IStruct interface
type StructArray []IStruct

func (r StructArray) ValidateType(searchedValue interface{}) error {
	if len(r) == 0 {
		log.Println("Empty slice has been found")
		return errors.New("slice is empty")
	}

	if err := r[0].CompareTypes(searchedValue); err != nil {
		return err
	}

	return nil
}

func (r StructArray) IsPresent(searchedValue interface{}) (bool, error) {
	if err := r.ValidateType(searchedValue); err != nil {
		return false, err
	}

	for index := range r {
		output, err := r[index].CompareValues(searchedValue)

		if err != nil {
			return false, err
		}

		if output {
			return true, nil
		}
	}

	return false, nil
}

func (r StructArray) FindIndex(searchedValue interface{}) (int, error) {
	if err := r.ValidateType(searchedValue); err != nil {
		return 0, err
	}

	for index := range r {
		output, err := r[index].CompareValues(searchedValue)

		if err != nil {
			return 0, err
		}

		if output {
			return index, nil
		}
	}

	return 0, nil
}

func (r StructArray) Get(searchedValue interface{}) (interface{}, error) {
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
