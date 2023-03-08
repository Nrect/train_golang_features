package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

type MatrixBody struct {
	Value  string // Буква
	Col    int32  // Номер колонки
	Row    string // Номер строки
	Cypher string // Шифр строка+колонка
}

func main() {
	words, clearString := GetDictionaryWords()
	fmt.Println(words)
	fmt.Println(clearString)

	uniqSymbols := make(map[string]string)
	cols := []string{"1", "2", "3", "4"}
	rows := []string{"А", "Б", "В", "Г", "Д", "Е", "Ж", "З"}

	matrixExercise := make(map[string]*MatrixBody)

	for i, w := 0, 0; i < len(clearString); i += w {
		runeValue, width := utf8.DecodeRuneInString(clearString[i:])

		convertedValue := string(runeValue)

		cypher := GenerateCypher(rows, cols)

		_, ok := uniqSymbols[convertedValue]
		if !ok {
			uniqSymbols[convertedValue] = convertedValue

			_, ok = matrixExercise[cypher]
			if !ok {
				matrixExercise[cypher] = nil
			} else {
				isOk := 0
				for isOk < 1 {
					newCypher := GenerateCypher(rows, cols)
					_, ok = matrixExercise[newCypher]
					if !ok {
						matrixExercise[newCypher] = nil
						isOk = 1
					}
				}
			}
		}

		w = width
	}

	fmt.Println(len(uniqSymbols))
	fmt.Println(len(matrixExercise))

}

func GenerateCypher(rows, cols []string) string {
	randomCol := rand.Intn(len(cols) - 0)
	randomRow := rand.Intn(len(rows) - 0)
	return rows[randomRow] + cols[randomCol]
}

func GetDictionaryWords() ([]string, string) {
	template := "весело;воробей;ворона;девочка;дежурный;деревня;заяц;карандаш;класс;корова;лисица;мальчик;машина;медведь;молоко"
	//template := "молоко"

	template = strings.TrimSpace(template)

	words := strings.Split(template, ";")
	clearString := strings.Replace(template, ";", "", -1)

	return words, clearString
}
