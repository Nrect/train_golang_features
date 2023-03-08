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
	assert.Equal(t, len(games.GetRows(1)), 1, "should be equal")
	assert.Equal(t, len(games.GetRows(4)), 1, "should be equal")
	assert.Equal(t, len(games.GetRows(8)), 2, "should be equal")
	assert.Equal(t, len(games.GetRows(12)), 3, "should be equal")
	assert.Equal(t, len(games.GetRows(16)), 4, "should be equal")
	assert.Equal(t, len(games.GetRows(20)), 5, "should be equal")
	assert.Equal(t, len(games.GetRows(24)), 6, "should be equal")
	assert.Equal(t, len(games.GetRows(28)), 7, "should be equal")
	assert.Equal(t, len(games.GetRows(32)), 8, "should be equal")

	assert.Equal(t, len(games.GetRows(-1)), 0, "should be equal")
	assert.Equal(t, len(games.GetRows(0)), 0, "should be equal")
	assert.Equal(t, len(games.GetRows(100)), 0, "should be equal")
}
