package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	inputString = []string{"bench", "curl", "dip"}
)

func TestDataInArray(t *testing.T) {
	assert.True(t, IsValueInSlice(inputString, "bench"))
}

func TestDataNotInArray(t *testing.T) {
	assert.False(t, IsValueInSlice(inputString, "hamstring"))
}

func TestDataInBothArraysStrict(t *testing.T) {
	searchedValues := []string{"dip", "bench"}
	assert.True(t, AreValuesInSliceStrict(inputString, searchedValues, true))
}

func TestDataInBothArraysStrictNotFound(t *testing.T) {
	searchedValues := []string{"foo", "bench"}
	assert.False(t, AreValuesInSliceStrict(inputString, searchedValues, true))
}

func TestDataInBothArraysNotStrict(t *testing.T) {
	searchedValues := []string{"foo", "bench"}
	assert.True(t, AreValuesInSliceStrict(inputString, searchedValues, false))
}

func TestDataInBothArraysNotStrictNotFound(t *testing.T) {
	searchedValues := []string{"foo", "bar"}
	assert.False(t, AreValuesInSliceStrict(inputString, searchedValues, false))
}
