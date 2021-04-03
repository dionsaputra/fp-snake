package math

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

func ternary(condition bool, positive int, negative int) int {
	if condition {
		return positive
	}
	return negative
}

func lazyTernary(condition bool, onPositive func() int, onNegative func() int) int {
	if condition {
		return onPositive()
	}
	return onNegative()
}

func Abs(x int) int {
	return ternary(x < 0, -x, x)
}

func Sign(x int) int {
	return lazyTernary(x == 0, func() int { return 0 }, func() int { return x / Abs(x) })
}
