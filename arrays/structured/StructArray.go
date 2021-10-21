package structured

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays"
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

func (r StructArray) GetByIndex(index int) (interface{}, error) {
	if index > len(r) || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return 0, errors.New(arrays.OutOfBoundsError)
	} else {
		return r[index], nil
	}
}

func (r *StructArray) Add(newValue interface{}) error {
	compatibleObj, ok := newValue.(IStruct)
	if !ok {
		log.Println("value is not of type IStruct")
		return errors.New(arrays.MismatchedTypeError)
	}

	*r = append(*r, compatibleObj)
	return nil
}

func (r *StructArray) RemoveByIndex(index int) error {
	sliceLength := len(*r)

	if index > sliceLength || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return errors.New(arrays.OutOfBoundsError)
	}

	tmp := make([]IStruct, sliceLength)
	copy(tmp, *r)

	tmp[index] = tmp[len(tmp)-1]
	*r = tmp[:len(tmp)-1]

	return nil
}
