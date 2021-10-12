package structured

import (
	"errors"
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestStructure struct {
	Name   string
	Number int
}

func (r TestStructure) CompareValues(input interface{}) (bool, error) {
	if err := r.CompareTypes(input); err != nil {
		return false, err
	}

	return r.Name == input, nil
}

func (r TestStructure) CompareTypes(input interface{}) error {
	arrayType := reflect.TypeOf(r.Name)
	paramType := reflect.TypeOf(input)

	if paramType.Kind() != arrayType.Kind() {
		return errors.New(arrays.MismatchedTypeError)
	} else {
		return nil
	}
}

var testArray = StructArray{
	TestStructure{
		Name:   "Foo",
		Number: 12345,
	},
	TestStructure{
		Name:   "Bar",
		Number: 23456,
	},
	TestStructure{
		Name:   "Xyz",
		Number: 34567,
	},
}

func Test_StructArrayOccurrence(t *testing.T) {
	output, err := testArray.IsPresent("Foo")
	assert.Nil(t, err)
	assert.True(t, output)
}

func Test_StructArrayNotFound(t *testing.T) {
	output, err := testArray.IsPresent("Out")
	assert.Nil(t, err)
	assert.False(t, output)
}

func Test_StructArrayTypeMismatch(t *testing.T) {
	_, err := testArray.IsPresent(123)
	assert.Error(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_StructArrayIndex(t *testing.T) {
	output, err := testArray.FindIndex("Xyz")
	assert.Nil(t, err)
	assert.Equal(t, 2, output)
}

func Test_StructArrayIndexNotFound(t *testing.T) {
	output, err := testArray.FindIndex("Out")
	assert.Nil(t, err)
	assert.Equal(t, 0, output)
}

func Test_StructArrayIndexTypeMismatch(t *testing.T) {
	_, err := testArray.IsPresent(123)
	assert.Error(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_StructGet(t *testing.T) {
	output, err := testArray.Get("Xyz")
	assert.Nil(t, err)
	assert.Equal(t, "Xyz", output.(TestStructure).Name)
	assert.Equal(t, 34567, output.(TestStructure).Number)
}

func Test_StructGetNotFound(t *testing.T) {
	_, err := testArray.Get("Out")
	assert.Error(t, err)
	assert.Equal(t, "could not find searched structure", err.Error())
}

func Test_StructGetTypeMismatch(t *testing.T) {
	_, err := testArray.Get(12345)
	assert.Error(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}
