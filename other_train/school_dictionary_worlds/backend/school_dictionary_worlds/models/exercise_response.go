package models

type ExerciseResponse struct {
	Matrix   map[string]*MatrixBody `json:"matrix"`
	Exercise []*ExerciseBody        `json:"exercise"`
}
