package main

import (
	"github.com/dionsaputra/fp-snake/deques"
	"github.com/dionsaputra/fp-snake/geometry"
	"github.com/dionsaputra/fp-snake/snakes"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := snakes.SnakeOf(geometry.PointOf(10, 10)).
		WithTail(deques.DequeOf(geometry.PointOf(10, 11), geometry.PointOf(10, 12)))

	width := 30
	height := 15
	for t := 0; t < 10; t++ {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				contains := s.Contains(geometry.PointOf(j, i))
				if contains {
					print("#")
				} else {
					print(".")
				}
				print(" ")
			}
			print("\n")
		}
		s = s.Eat()
		time.Sleep(300 * 1000 * 1000)
	}
}
