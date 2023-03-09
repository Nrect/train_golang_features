package test

import (
	"github.com/stretchr/testify/assert"
	"school_dictionary_worlds/games"
	"testing"
)

const TEMPLATE = "библиотека"

func TestFillMatrix(t *testing.T) {
	_, clearString := games.GetDictionaryWords()

	uniqSymbols := games.GenerateUniqSymbols(clearString)
	matrixExercise := games.GenerateMatrixExercise(uniqSymbols)
	//TODO скоро этот тест не будет валидным, тк придется матрицу заполнять случайными значениями
	assert.Equal(t, len(uniqSymbols), len(matrixExercise), "len uniqSymbols and matrixExercise should be equal")
}

func TestGenerateUniqSymbols(t *testing.T) {
	result := games.GenerateUniqSymbols(TEMPLATE)
	need := map[string]string{"б": "б", "и": "и", "л": "л", "о": "о", "т": "т", "е": "е", "к": "к", "а": "а"}

	assert.Equal(t, need, result, "should be equal")
}

func TestGenerateCypher(t *testing.T) {
	cols := games.GetCols()
	rows := games.GetRows(1)
	cypher, col, row := games.GenerateCypher(rows, cols)

	assert.NotNil(t, cypher, "should not be nil")
	assert.NotNil(t, col, "should not be nil")
	assert.NotNil(t, row, "should not be nil")
}

func TestGetCols(t *testing.T) {
	cols := games.GetCols()
	testData := []string{"1", "2", "3", "4"}
	assert.Equal(t, cols, testData, "cols and testData should be equal")
}

func TestGetRows(t *testing.T) {
	assert.Equal(t, 1, len(games.GetRows(1)), "should be equal")
	assert.Equal(t, 1, len(games.GetRows(4)), "should be equal")
	assert.Equal(t, 2, len(games.GetRows(8)), "should be equal")
	assert.Equal(t, 3, len(games.GetRows(12)), "should be equal")
	assert.Equal(t, 4, len(games.GetRows(16)), "should be equal")
	assert.Equal(t, 5, len(games.GetRows(20)), "should be equal")
	assert.Equal(t, 6, len(games.GetRows(24)), "should be equal")
	assert.Equal(t, 7, len(games.GetRows(28)), "should be equal")
	assert.Equal(t, 8, len(games.GetRows(32)), "should be equal")

	assert.Equal(t, 0, len(games.GetRows(-1)), "should be equal")
	assert.Equal(t, 0, len(games.GetRows(0)), "should be equal")
	assert.Equal(t, 0, len(games.GetRows(100)), "should be equal")
}

func TestGetDictionaryWords(t *testing.T) {
	words, clearString := games.GetDictionaryWords()

	assert.NotNil(t, words, "should not be nil")
	assert.NotNil(t, clearString, "should not be nil")
}
