package models

type ExerciseResponse struct {
	Matrix   map[string]*MatrixBody `json:"matrix"`
	Exercise []*ExerciseBody        `json:"exercises"`
	Rows     []string               `json:"rows"`
}
