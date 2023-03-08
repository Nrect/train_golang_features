package games

import (
	"fmt"
	"math/rand"
	"school_dictionary_worlds/models"
	"strings"
	"unicode/utf8"
)

func LunchCryptographer() {
	words, clearString := GetDictionaryWords()
	fmt.Println(words)

	uniqSymbols, matrixExercise := FillMatrix(clearString)
	fmt.Println(uniqSymbols)
	fmt.Println(matrixExercise)
}

// FillMatrix fill the matrix with unique elements from clear string. Return uniqSymbols and matrixExercise
func FillMatrix(clearString string) (map[string]string, map[string]*models.MatrixBody) {
	uniqSymbols := make(map[string]string)
	matrixExercise := make(map[string]*models.MatrixBody)

	cols := GetCols()
	rows := GetRows(len(uniqSymbols))

	for i, w := 0, 0; i < len(clearString); i += w {
		runeValue, width := utf8.DecodeRuneInString(clearString[i:])

		convertedValue := string(runeValue)

		_, ok := uniqSymbols[convertedValue]
		if !ok {
			uniqSymbols[convertedValue] = convertedValue
		}

		w = width
	}

	for _, symbol := range uniqSymbols {
		cypher, col, row := GenerateCypher(rows, cols)

		_, ok := matrixExercise[cypher]
		if !ok {
			matrixExercise[cypher] = &models.MatrixBody{
				Value:  symbol,
				Col:    col,
				Row:    row,
				Cypher: cypher,
			}
		} else {
			isOk := 0
			for isOk < 1 {
				newCypher, col1, row1 := GenerateCypher(rows, cols)
				_, ok = matrixExercise[newCypher]
				if !ok {
					matrixExercise[newCypher] = &models.MatrixBody{
						Value:  symbol,
						Col:    col1,
						Row:    row1,
						Cypher: cypher,
					}
					isOk = 1
				}
			}
		}
	}

	return uniqSymbols, matrixExercise
}

func GetCols() []string {
	return []string{"1", "2", "3", "4"}
}

func GetRows(uniqSymbolsCount int) []string {
	return []string{"А", "Б", "В", "Г", "Д", "Е", "Ж", "З"}
}

func GenerateCypher(rows, cols []string) (string, int, int) {
	randomCol := rand.Intn(len(cols) - 0)
	randomRow := rand.Intn(len(rows) - 0)
	return rows[randomRow] + cols[randomCol], randomCol, randomRow
}

func GetDictionaryWords() ([]string, string) {
	template := "весело;воробей;ворона;девочка;дежурный;деревня;заяц;карандаш;класс;корова;лисица;мальчик;машина;медведь;молоко"
	//template := "молоко"

	template = strings.TrimSpace(template)

	words := strings.Split(template, ";")
	clearString := strings.Replace(template, ";", "", -1)

	return words, clearString
}
