package string

import (
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStringArray = StringArray{"Foo", "Bar", "Xyz"}

func Test_StringArrayOccurrence(t *testing.T) {
	output, err := testStringArray.IsPresent("Foo")
	assert.Nil(t, err)
	assert.True(t, output)
}

func Test_StringArrayOccurrenceMismatchedType(t *testing.T) {
	output, err := testStringArray.IsPresent(0)
	assert.Error(t, err)
	assert.False(t, output)
}

func Test_StringArrayOccurrenceNotFound(t *testing.T) {
	output, err := testStringArray.IsPresent("Out")
	assert.Nil(t, err)
	assert.False(t, output)
}

func Test_StringArrayFoundIndex(t *testing.T) {
	output, err := testStringArray.FindIndex("Bar")
	assert.Nil(t, err)
	assert.Equal(t, 1, output)
}

func Test_StringArrayFoundIndexMismatchType(t *testing.T) {
	output, err := testStringArray.FindIndex("Bar")
	assert.Nil(t, err)
	assert.Equal(t, 1, output)
}

func Test_StringArrayIndexZeroResults(t *testing.T) {
	output, err := testStringArray.FindIndex("Out")

	assert.NotNil(t, err)
	assert.Equal(t, arrays.NotFoundError, err.Error())
	assert.Equal(t, -1, output)
}

func Test_GetByFirstIndex(t *testing.T) {
	output, err := testStringArray.GetByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", output)
}

func Test_GetByIndex(t *testing.T) {
	output, err := testStringArray.GetByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "Bar", output)
}

func Test_GetByLastIndex(t *testing.T) {
	output, err := testStringArray.GetByIndex(len(testStringArray) - 1)
	assert.Nil(t, err)
	assert.Equal(t, "Xyz", output)
}

func Test_GetByIndexNotFound(t *testing.T) {
	_, err := testStringArray.GetByIndex(10)
	assert.Error(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}

func Test_AddToSlice(t *testing.T) {
	newValue := "NewOne"

	err := testStringArray.Add(newValue)
	assert.Nil(t, err)

	isPresent, err := testStringArray.IsPresent(newValue)
	assert.Nil(t, err)
	assert.True(t, isPresent)
}

func Test_AddToSliceIncompatible(t *testing.T) {
	newValue := 55
	err := testStringArray.Add(newValue)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_RemoveFromSlice(t *testing.T) {
	err := testStringArray.RemoveByIndex(1)
	assert.Nil(t, err)

	isPresent, err := testStringArray.IsPresent("Bar")
	assert.Nil(t, err)
	assert.False(t, isPresent)
}

func Test_RemoveFromSliceOutOfBounds(t *testing.T) {
	err := testStringArray.RemoveByIndex(10)
	assert.NotNil(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}
