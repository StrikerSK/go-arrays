package string

import (
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/StrikerSK/go-arrays/arrays/exception"
	"log"
	"reflect"
)

type StringArray []string

func (r StringArray) FindIndex(searchedValue interface{}) (int, error) {
	if err := arrays.CheckExpectedType(searchedValue, reflect.String); err != nil {
		return -1, err
	}

	for index := range r {
		if r[index] == searchedValue {
			return index, nil
		}
	}

	return -1, exception.NewNotFoundException()
}

func (r StringArray) IsPresent(searchedValue interface{}) (bool, error) {
	index, err := r.FindIndex(searchedValue)

	if err != nil && err.Error() != exception.NotFoundError {
		return false, err
	}

	return index >= 0, nil
}

func (r StringArray) GetByIndex(index int) (interface{}, error) {
	if index > len(r) || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return 0, exception.NewOutOfBoundsException()
	} else {
		return r[index], nil
	}
}

func (r *StringArray) Add(newValue interface{}) error {
	if err := arrays.CheckExpectedType(newValue, reflect.String); err != nil {
		return err
	}

	*r = append(*r, newValue.(string))
	return nil
}

func (r *StringArray) RemoveByIndex(index int) error {
	sliceLength := len(*r)

	if index > sliceLength || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return exception.NewOutOfBoundsException()
	}

	tmp := make([]string, sliceLength)
	copy(tmp, *r)

	tmp[index] = tmp[len(tmp)-1]
	*r = tmp[:len(tmp)-1]

	return nil
}
