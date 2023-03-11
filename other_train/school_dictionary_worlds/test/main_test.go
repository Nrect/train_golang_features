package test

import (
	"github.com/stretchr/testify/assert"
	"school_dictionary_worlds/games"
	"school_dictionary_worlds/models"
	"testing"
)

const Template = "библиотека"

func GetTemplateMap() map[string]string {
	return map[string]string{"б": "б", "и": "и", "л": "л", "о": "о", "т": "т", "е": "е", "к": "к", "а": "а"}
}

func TestGenerateExercise(t *testing.T) {
	words := []string{Template, Template}
	matrixExercise := map[string]*models.MatrixBody{
		"А1": {
			Value:  "б",
			Cypher: "А1",
		},
		"А2": {
			Value:  "и",
			Cypher: "А2",
		},
		"А3": {
			Value:  "л",
			Cypher: "А3",
		},
		"А4": {
			Value:  "о",
			Cypher: "А4",
		},
		"Б1": {
			Value:  "т",
			Cypher: "Б1",
		},
		"Б2": {
			Value:  "е",
			Cypher: "Б2",
		},
		"Б3": {
			Value:  "к",
			Cypher: "Б3",
		},
		"Б4": {
			Value:  "а",
			Cypher: "Б4",
		},
	}

	result := games.GenerateExercise(words, matrixExercise)
	need := []*models.ExerciseBody{
		{
			Exercise: "А1 А2 А1 А3 А2 А4 Б1 Б2 Б3 Б4",
			Answer:   Template,
		},
		{
			Exercise: "А1 А2 А1 А3 А2 А4 Б1 Б2 Б3 Б4",
			Answer:   Template,
		},
	}

	assert.Equal(t, need[0].Exercise, result[0].Exercise, "Exercise should be equal")
	assert.Equal(t, need[0].Answer, result[0].Answer, "Answer should be equal")
}

func TestFillMatrix(t *testing.T) {
	_, clearString := games.GetDictionaryWords()

	uniqSymbols := games.GenerateUniqSymbols(clearString)
	matrixExercise := games.GenerateMatrixExercise(uniqSymbols)
	//TODO скоро этот тест не будет валидным, тк придется матрицу заполнять случайными значениями
	assert.Equal(t, len(uniqSymbols), len(matrixExercise), "len uniqSymbols and matrixExercise should be equal")
}

func TestGenerateUniqSymbols(t *testing.T) {
	result := games.GenerateUniqSymbols(Template)
	need := GetTemplateMap()

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
