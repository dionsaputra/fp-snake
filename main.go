package main

import (
	"github.com/dionsaputra/fp-snake/logics"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := logics.NewSnake(logics.NewSegment(5, 5)).
		Grow(logics.NewDirection(1, 0), logics.NewDimension(30, 15)).
		Grow(logics.NewDirection(1, 0), logics.NewDimension(30, 15))

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
		s = s.Move(logics.NewDirection(1, 0), logics.NewDimension(30, 15))
		time.Sleep(300 * 1000 * 1000)
	}
}
