package test

import (
	"github.com/stretchr/testify/assert"
	"school_dictionary_worlds/games"
	"testing"
)

func TestFillMatrix(t *testing.T) {
	_, clearString := games.GetDictionaryWords()
	uniqSymbols, matrixExercise := games.FillMatrix(clearString)
	assert.Equal(t, len(uniqSymbols), len(matrixExercise), "len uniqSymbols and matrixExercise should be equal")
}

func TestGetCols(t *testing.T) {
	cols := games.GetCols()
	testData := []string{"1", "2", "3", "4"}
	assert.Equal(t, cols, testData, "cols and testData should be equal")
}

func TestGetRows(t *testing.T) {
	assert.Equal(t, len(games.GetRows(8)), 8, "should be equal")
}
