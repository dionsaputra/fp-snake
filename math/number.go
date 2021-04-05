package math

func Mod(num int, div int) int {
	return (num%div + div) % div
}
