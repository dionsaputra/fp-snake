package main

import (
	"github.com/dionsaputra/fp-snake/math"
	"github.com/dionsaputra/fp-snake/snake"
	"os"
	"os/exec"
	"time"
)

func main() {
	dimension := math.Dimension{Height: 15, Width: 30}
	right := math.Right()

	s := snake.Snake{
		Head: snake.Head{
			Segment:   snake.Segment{Row: 5, Col: 5},
			Direction: right,
		},
	}
	s = s.Grow(dimension).Grow(dimension)

	width := 30
	height := 15
	for t := 0; t < 100; t++ {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				contains := s.Contains(snake.Segment{Row: i, Col: j})
				if contains {
					print("#")
				} else {
					print(".")
				}
				print(" ")
			}
			print("\n")
		}
		s = s.Move(dimension)
		time.Sleep(300 * 1000 * 1000)
	}
}
