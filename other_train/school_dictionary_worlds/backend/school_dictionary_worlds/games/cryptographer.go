package games

import (
	"math/rand"
	"school_dictionary_worlds/models"
	"strings"
	"unicode/utf8"
)

func LunchCryptographer() *models.ExerciseResponse {
	words, clearString := GetDictionaryWords()

	uniqSymbols := GenerateUniqSymbols(clearString)
	matrixExercise := GenerateMatrixExercise(uniqSymbols)

	exercise := GenerateExercise(words, matrixExercise)

	response := models.ExerciseResponse{
		Matrix:   matrixExercise,
		Exercise: exercise,
	}

	return &response
}

// TODO: УБРАТЬ ЛОГИКУ КОДИРОВОК ТК ГОШКА И С БИТАМИ НОРМ РАБОТАЕТ)

// TODO: сделать логику, если у нас нечетное кол-во, то добавить в пустые ячейки случайные символы
// Опять же, если символ прям случайный, если нет, то какой эффективный алгоритм поиска другого символа?
// Сделать строку с ру словарем и строку с уник символами и сравнить их разницу
// Для начала сделаю со случайным значением, затем уже и с уникальным

// СЕЙЧАС ПОРЯДОК НАПИСАНИЯ ФУНКЦИЙ ЗАВИСТ ОТ ИХ ИСПОЛЬЗОВАНЯ СВЕРХУ ВНИЗ

// GenerateExercise generate matrix based exercise
func GenerateExercise(words []string, matrixExercise map[string]*models.MatrixBody) []*models.ExerciseBody {

	result := make([]*models.ExerciseBody, 0, len(words)-1)

	for _, word := range words {
		exercise := ""

		for _, character := range word {

			for _, m := range matrixExercise {
				if m.Value == string(character) {
					exercise += m.Cypher + " "
				}
			}

		}

		result = append(result, &models.ExerciseBody{
			Exercise: strings.TrimSpace(exercise),
			Answer:   word,
		})

	}

	return result
}

// GenerateMatrixExercise fill the matrix with unique elements. Return uniqSymbols and matrixExercise
func GenerateMatrixExercise(uniqSymbols map[string]string) map[string]*models.MatrixBody {
	matrixExercise := make(map[string]*models.MatrixBody)

	cols := GetCols()
	rows := GetRows(len(uniqSymbols))

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

	return matrixExercise
}

// GenerateUniqSymbols generate unique symbols map from clearString
func GenerateUniqSymbols(clearString string) map[string]string {
	uniqSymbols := make(map[string]string)

	for i, w := 0, 0; i < len(clearString); i += w {
		runeValue, width := utf8.DecodeRuneInString(clearString[i:])

		convertedValue := string(runeValue)

		_, ok := uniqSymbols[convertedValue]
		if !ok {
			uniqSymbols[convertedValue] = convertedValue
		}

		w = width
	}

	return uniqSymbols
}

// GenerateCypher generate data for matrix exercise
func GenerateCypher(rows, cols []string) (string, int, int) {
	randomCol := rand.Intn(len(cols) - 0)
	randomRow := rand.Intn(len(rows) - 0)
	return rows[randomRow] + cols[randomCol], randomCol, randomRow
}

// GetCols get static 4 columns
func GetCols() []string {
	return []string{"1", "2", "3", "4"}
}

// GetRows get dynamic rows. Length row depend on uniqSymbolsCount
func GetRows(uniqSymbolsCount int) []string {
	template := []string{"А", "Б", "В", "Г", "Д", "Е", "Ж", "З"}

	if uniqSymbolsCount <= 0 || uniqSymbolsCount > 32 {
		template = template[:0]
	} else if uniqSymbolsCount <= 4 {
		template = template[:1]
	} else if uniqSymbolsCount <= 8 {
		template = template[:2]
	} else if uniqSymbolsCount <= 12 {
		template = template[:3]
	} else if uniqSymbolsCount <= 16 {
		template = template[:4]
	} else if uniqSymbolsCount <= 20 {
		template = template[:5]
	} else if uniqSymbolsCount <= 24 {
		template = template[:6]
	} else if uniqSymbolsCount <= 28 {
		template = template[:7]
	} else if uniqSymbolsCount <= 32 {
		template = template[:8]
	}

	return template
}

// GetDictionaryWords get data for building game
func GetDictionaryWords() ([]string, string) {
	template := "весело;воробей;ворона;девочка;дежурный;деревня;заяц;карандаш;класс;корова;лисица;мальчик;машина;медведь;молоко"
	//template := "молоко"

	template = strings.TrimSpace(template)

	words := strings.Split(template, ";")
	clearString := strings.Replace(template, ";", "", -1)

	return words[:5], clearString
}
