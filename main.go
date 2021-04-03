package main

import (
	"github.com/dionsaputra/fp-snake/geometry"
	"github.com/dionsaputra/fp-snake/snakes"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := snakes.SnakeOf(geometry.PointOf(10, 10))
	s.Tail().PushBack(geometry.PointOf(10, 11))
	s.Tail().PushBack(geometry.PointOf(10, 12))

	//fmt.Printf("%v\n", reflect.TypeOf(s.Tail()))
	//println(s.Tail().Size())
	//os.Exit(1)

	width := 30
	height := 15
	for t := 0; t < 10; t++ {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				contains := s.Contains(geometry.PointOf(j, i))
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
		s = s.Eat()
		time.Sleep(300 * 1000 * 1000)
	}
}
