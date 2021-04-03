package maths

import (
	"math/rand"
	"time"
)

func RandInt(start, end int) int {
	random := int(rand.NewSource(time.Now().Unix()).Int63())
	return start + (random % (end - start))
}

func Mod(num int, div int) int {
	return (num + div) % div
}