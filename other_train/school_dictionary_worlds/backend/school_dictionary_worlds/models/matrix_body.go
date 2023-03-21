package models

type MatrixBody struct {
	Value   string `json:"value"`  // Буква
	Col     int    `json:"col"`    // Номер колонки
	Row     int    `json:"row"`    // Номер строки
	Cypher  string `json:"cypher"` // Шифр строка+колонка
	RowName string `json:"rowName"`
}
