package main

import (
	"github.com/dionsaputra/fp-snake/points"
	"github.com/dionsaputra/fp-snake/snakes"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := snakes.SnakeOf(points.PointOf(5, 17))
	s.Tail().PushBack(points.PointOf(5, 18))
	s.Tail().PushBack(points.PointOf(5, 19))

	//fmt.Printf("%v\n", reflect.TypeOf(s.Tail()))
	//println(s.Tail().Size())
	//os.Exit(1)

	width := 50
	height := 20
	for t := 0; t < 15; t++ {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				contains := s.Contains(points.PointOf(i, j))
				//fmt.Println(i, j, contains)
				if contains {
					print("#")
				} else {
					print(".")
				}
				print(" ")
			}
			print("\n")
		}
		s = s.Move(1, 0)
		time.Sleep(300 * 1000 * 1000)
	}
}
