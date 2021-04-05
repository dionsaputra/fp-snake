package main

import (
	"github.com/dionsaputra/fp-snake/logics"
	"os"
	"os/exec"
	"time"
)

func main() {
	dimension := logics.NewDimension(30, 15)
	right := logics.NewDirection(1, 0)

	s := logics.NewSnake(logics.NewSegment(5, 5)).
		Grow(right, dimension).
		Grow(right, dimension)

	width := 30
	height := 15
	for t := 0; t < 100; t++ {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				contains := s.Contains(logics.NewSegment(j, i))
				if contains {
					print("#")
				} else {
					print(".")
				}
				print(" ")
			}
			print("\n")
		}
		s = s.Move(right, dimension)
		time.Sleep(300 * 1000 * 1000)
	}
}
