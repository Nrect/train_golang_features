package utils

import "math/rand"

func GetRandomMoneyValue() float64 {
	return rand.Float64() * 10000
}
