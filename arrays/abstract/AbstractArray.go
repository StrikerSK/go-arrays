package abstract

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays"
	"log"
	"reflect"
)

// AbstractArray - Defined for core types as string, int, int8, int16, int32,...
type AbstractArray []interface{}

func (r AbstractArray) ArrayType() string {
	return reflect.TypeOf(r[0]).Name()
}

func (r AbstractArray) ValidateType(searchedValue interface{}) error {
	if len(r) == 0 {
		log.Println("Empty slice has been found")
		return errors.New("slice is empty")
	}

	arrayType := reflect.TypeOf(r[0])
	paramType := reflect.TypeOf(searchedValue)
	isEqual := arrayType == paramType

	if paramType.Kind() == reflect.Struct || arrayType.Kind() == reflect.Struct {
		log.Println("Struct cannot be used")
		return errors.New(arrays.StructUseError)
	}

	if !isEqual {
		log.Printf("Array type [%s] should equal parameter type [%s]\n", arrayType, paramType)
		return errors.New(arrays.MismatchedTypeError)
	}

	return nil
}

func (r AbstractArray) IsPresent(searchedValue interface{}) (bool, error) {
	if err := r.ValidateType(searchedValue); err != nil {
		return false, err
	}

	for index := range r {
		if r[index] == searchedValue {
			return true, nil
		}
	}

	return false, nil
}

func (r AbstractArray) FindIndex(searchedValue interface{}) (int, error) {
	if err := r.ValidateType(searchedValue); err != nil {
		return 0, err
	}

	for index := range r {
		if r[index] == searchedValue {
			return index, nil
		}
	}

	return 0, nil
}

func (r AbstractArray) Get(searchedValue interface{}) (interface{}, error) {
	index, err := r.FindIndex(searchedValue)
	if err != nil {
		return nil, err
	}

	if index == 0 {
		return nil, errors.New("could not found searched structure")
	} else {
		return r[index], nil
	}
}
