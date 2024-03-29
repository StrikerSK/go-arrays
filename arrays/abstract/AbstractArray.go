package abstract

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays/exception"
	"log"
	"reflect"
)

// AbstractArray - defines array for core types such as string, int, int8, int16, int32,...
type AbstractArray []interface{}

// ArrayType - defined base on first element
func (r AbstractArray) ArrayType() string {
	if len(r) == 0 {
		return "empty"
	}

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
		return errors.New(exception.StructUseError)
	}

	if !isEqual {
		log.Printf("Array type [%s] should equal parameter type [%s]\n", arrayType, paramType)
		return exception.NewMismatchException()
	}

	return nil
}

func (r AbstractArray) FindIndex(searchedValue interface{}) (int, error) {
	if err := r.validateType(searchedValue); err != nil {
		return -1, err
	}

	for index := range r {
		if r[index] == searchedValue {
			return index, nil
		}
	}

	return -1, exception.NewNotFoundException()
}

func (r AbstractArray) IsPresent(searchedValue interface{}) (bool, error) {
	index, err := r.FindIndex(searchedValue)

	if err != nil && err.Error() != exception.NotFoundError {
		return false, err
	}

	return index >= 0, nil
}

func (r AbstractArray) GetByIndex(index int) (interface{}, error) {
	if index > len(r) || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return 0, exception.NewOutOfBoundsException()
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

func (r *AbstractArray) RemoveByIndex(index int) error {
	sliceLength := len(*r)

	if index > sliceLength || index < 0 {
		log.Println("Provided index parameter is out of bounds")
		return exception.NewOutOfBoundsException()
	}

	tmp := make([]interface{}, sliceLength)
	copy(tmp, *r)

	tmp[index] = tmp[len(tmp)-1]
	*r = tmp[:len(tmp)-1]

	return nil
}
