package structured

import (
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/stretchr/testify/assert"
	"log"
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

func (r TestStructure) CompareTypes(input interface{}) (err error) {
	err = arrays.ValidateArray(input, r.Name)
	return
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
	assert.NotNil(t, err)
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

func Test_GetByFirstIndex(t *testing.T) {
	output, err := testArray.GetByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", output.(TestStructure).Name)
	assert.Equal(t, 12345, output.(TestStructure).Number)
}

func Test_GetByIndex(t *testing.T) {
	output, err := testArray.GetByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "Bar", output.(TestStructure).Name)
	assert.Equal(t, 23456, output.(TestStructure).Number)
}

func Test_GetByLastIndex(t *testing.T) {
	output, err := testArray.GetByIndex(len(testArray) - 1)
	assert.Nil(t, err)
	assert.Equal(t, "Xyz", output.(TestStructure).Name)
	assert.Equal(t, 34567, output.(TestStructure).Number)
}

func Test_GetByIndexNotFound(t *testing.T) {
	_, err := testArray.GetByIndex(4)
	assert.Error(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}

func Test_AddToSlice(t *testing.T) {
	newValue := TestStructure{
		Name:   "NewValue",
		Number: 911,
	}

	err := testArray.Add(newValue)
	assert.Nil(t, err)

	isPresent, err := testArray.IsPresent(newValue.Name)
	assert.Nil(t, err)
	assert.True(t, isPresent)
}

func Test_AddToSliceIncompatible(t *testing.T) {
	err := testArray.Add("123")
	assert.NotNil(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_RemoveFromSlice(t *testing.T) {
	err := testArray.RemoveByIndex(1)
	assert.Nil(t, err)

	isPresent, err := testArray.IsPresent("Bar")
	assert.Nil(t, err)
	assert.False(t, isPresent)
}

func Test_RemoveFromSliceOutOfBounds(t *testing.T) {
	err := testArray.RemoveByIndex(10)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}

func Test_RUN(t *testing.T) {
	log.Println(reflect.TypeOf([]string{"Hello", "World"}).Elem())
}
