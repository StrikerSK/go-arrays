package abstract

import (
	"github.com/StrikerSK/go-arrays/arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testAbstractArray = AbstractArray{"Foo", "Bar", "Xyz"}

type testStruct struct {
	Name string
}

func Test_AssertArrayType(t *testing.T) {
	output := testAbstractArray.ArrayType()
	assert.Equal(t, "string", output)
}

func Test_AbstractEqualType(t *testing.T) {
	assert.Nil(t, testAbstractArray.validateType("Hello"))
}

func Test_AbstractStructParamType(t *testing.T) {
	err := testAbstractArray.validateType(testStruct{Name: "Foo"})
	assert.Error(t, err)
	assert.Equal(t, arrays.StructUseError, err.Error())
}

func Test_AbstractStructArrayType(t *testing.T) {
	s := AbstractArray{testStruct{Name: "Foo"}, testStruct{Name: "Bar"}}
	err := s.validateType("Foo")
	assert.Error(t, err)
	assert.Equal(t, arrays.StructUseError, err.Error())
}

func Test_AbstractNonEqualType(t *testing.T) {
	err := testAbstractArray.validateType(5)
	assert.Error(t, err)
	assert.Equal(t, arrays.MismatchedTypeError, err.Error())
}

func Test_AbstractArraySearchElementFound(t *testing.T) {
	output, err := testAbstractArray.IsPresent("Foo")
	assert.Nil(t, err)
	assert.True(t, output)
}

func Test_AbstractArraySearchElementNotFound(t *testing.T) {
	output, err := testAbstractArray.IsPresent("Hello")
	assert.Nil(t, err)
	assert.False(t, output)
}

func Test_AbstractArraySearchIncompatibleType(t *testing.T) {
	output, err := testAbstractArray.IsPresent(999)
	assert.Error(t, err)
	assert.False(t, output)
}

func Test_AbstractArrayIndexSearch(t *testing.T) {
	output, err := testAbstractArray.FindIndex("Bar")
	assert.Nil(t, err)
	assert.Equal(t, 1, output)
}

func Test_StringArrayIndexSearchNotFound(t *testing.T) {
	output, err := testAbstractArray.FindIndex("Hello")
	assert.Nil(t, err)
	assert.Equal(t, 0, output)
}

func Test_StringArrayIndexSearchIncompatibleType(t *testing.T) {
	output, err := testAbstractArray.FindIndex(999)
	assert.Error(t, err)
	assert.Equal(t, 0, output)
}

func Test_GetByFirstIndex(t *testing.T) {
	output, err := testAbstractArray.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", output)
}

func Test_GetByIndex(t *testing.T) {
	output, err := testAbstractArray.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, "Bar", output)
}

func Test_GetByLastIndex(t *testing.T) {
	output, err := testAbstractArray.Get(len(testAbstractArray) - 1)
	assert.Nil(t, err)
	assert.Equal(t, "Xyz", output)
}

func Test_GetByIndexNotFound(t *testing.T) {
	_, err := testAbstractArray.Get(10)
	assert.Error(t, err)
	assert.Equal(t, arrays.OutOfBoundsError, err.Error())
}
