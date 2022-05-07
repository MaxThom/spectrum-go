package utils

import "math/rand"

func RandomInt(start, limit int) int {
	return rand.Intn(limit) + start
}
