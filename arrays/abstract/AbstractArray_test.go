package abstract

import (
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stringArray = AbstractArray{"Foo", "Bar", "Xyz"}

type testStruct struct {
	Name string
}

func Test_AssertCorrectArrayType(t *testing.T) {
	assert.Equal(t, "string", stringArray.ArrayType())
	assert.Equal(t, "int", AbstractArray{1, 2, 3, 4}.ArrayType())
	assert.Equal(t, "empty", AbstractArray{}.ArrayType())
}

func Test_AbstractEqualType(t *testing.T) {
	assert.Nil(t, stringArray.validateType("Hello"))
}

func Test_AbstractStructParamType(t *testing.T) {
	err := stringArray.validateType(testStruct{Name: "Foo"})
	assert.NotNil(t, err)
	assert.Equal(t, arrays.StructUseError, err.Error())
}

func Test_AbstractStructArrayType(t *testing.T) {
	s := AbstractArray{testStruct{Name: "Foo"}, testStruct{Name: "Bar"}}
	err := s.validateType("Foo")
	assert.NotNil(t, err)
	assert.Equal(t, arrays.StructUseError, err.Error())
}

func Test_AbstractNonEqualType(t *testing.T) {
	err := stringArray.validateType(5)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_AbstractArraySearchElementFound(t *testing.T) {
	output, err := stringArray.IsPresent("Foo")
	assert.Nil(t, err)
	assert.True(t, output)
}

func Test_AbstractArraySearchElementNotFound(t *testing.T) {
	output, err := stringArray.IsPresent("Hello")
	assert.Nil(t, err)
	assert.False(t, output)
}

func Test_AbstractArraySearchIncompatibleType(t *testing.T) {
	output, err := stringArray.IsPresent(999)
	assert.Error(t, err)
	assert.False(t, output)
}

func Test_AbstractArrayIndexSearch(t *testing.T) {
	output, err := stringArray.FindIndex("Bar")
	assert.Nil(t, err)
	assert.Equal(t, 1, output)
}

func Test_StringArrayIndexSearchNotFound(t *testing.T) {
	output, err := stringArray.FindIndex("Hello")

	assert.NotNil(t, err)
	assert.Equal(t, arrays.NotFoundError, err.Error())
	assert.Equal(t, -1, output)
}

func Test_StringArrayIndexSearchIncompatibleType(t *testing.T) {
	output, err := stringArray.FindIndex(999)
	assert.Error(t, err)
	assert.Equal(t, -1, output)
}

func Test_GetByFirstIndex(t *testing.T) {
	output, err := stringArray.GetByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", output)
}

func Test_GetByIndex(t *testing.T) {
	output, err := stringArray.GetByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "Bar", output)
}

func Test_GetByLastIndex(t *testing.T) {
	output, err := stringArray.GetByIndex(len(stringArray) - 1)
	assert.Nil(t, err)
	assert.Equal(t, "Xyz", output)
}

func Test_GetByIndexNotFound(t *testing.T) {
	_, err := stringArray.GetByIndex(10)
	assert.Error(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}

func Test_AddToSlice(t *testing.T) {
	newValue := "NewOne"

	err := stringArray.Add(newValue)
	assert.Nil(t, err)

	isPresent, err := stringArray.IsPresent(newValue)
	assert.Nil(t, err)
	assert.True(t, isPresent)
}

func Test_AddToSliceIncompatible(t *testing.T) {
	newValue := 55
	err := stringArray.Add(newValue)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_RemoveFromSlice(t *testing.T) {
	err := stringArray.RemoveByIndex(1)
	assert.Nil(t, err)

	isPresent, err := stringArray.IsPresent("Bar")
	assert.Nil(t, err)
	assert.False(t, isPresent)
}

func Test_RemoveFromSliceOutOfBounds(t *testing.T) {
	err := stringArray.RemoveByIndex(10)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}
