package abstract

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays"
	"log"
	"reflect"
)

// AbstractArray - defines array for core types such as string, int, int8, int16, int32,...
type AbstractArray []interface{}

// ArrayType - defined base on first element
func (r AbstractArray) ArrayType() string {
	return reflect.TypeOf(r[0]).Name()
}

func (r AbstractArray) validateType(searchedValue interface{}) error {
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
	if err := r.validateType(searchedValue); err != nil {
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
	if err := r.validateType(searchedValue); err != nil {
		return 0, err
	}

	for index := range r {
		if r[index] == searchedValue {
			return index, nil
		}
	}

	return 0, nil
}

func (r AbstractArray) Get(index int) (interface{}, error) {
	if index > len(r) || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return 0, errors.New(arrays.OutOfBoundsError)
	} else {
		return r[index], nil
	}
}

func (r *AbstractArray) Add(newValue interface{}) error {
	if err := r.validateType(newValue); err != nil {
		return err
	}
	*r = append(*r, newValue)
	return nil
}
