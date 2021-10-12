package string

import (
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
	assert.Nil(t, err)
	assert.Equal(t, 0, output)
}
