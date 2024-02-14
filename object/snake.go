package object

import (
	"fmt"
	term "github.com/nsf/termbox-go"
)

type Snake struct {
	point *Point
}

func (s *Snake) GetPoints() Points {
	return []*Point{s.point}
}

func (s *Snake) Change(point *Point) {
	s.point = point
}

func (s *Snake) move(listener chan<- bool) {

	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				term.Sync()
				fmt.Println("ESC pressed")
				panic("DIE")
			case term.KeyArrowRight:
				term.Sync()
				fmt.Println("Arrow Right pressed")
				s.Change(NewPoint(s.point.X+1, s.point.Y))
			case term.KeyArrowLeft:
				term.Sync()
				fmt.Println("Arrow Left pressed")
				s.Change(NewPoint(s.point.X-1, s.point.Y))
			case term.KeyArrowUp:
				term.Sync()
				s.Change(NewPoint(s.point.X, s.point.Y+1))
				fmt.Println("Arrow Up pressed")
			case term.KeyArrowDown:
				term.Sync()
				s.Change(NewPoint(s.point.X, s.point.Y-1))
				fmt.Println("Arrow Down pressed")

			default:
				term.Sync()
				fmt.Println("ASCII : ", ev.Ch)

			}
			listener <- true
		case term.EventError:
			panic(ev.Err)
		}
	}
	//return nil
}

func NewSnake(point *Point, printListener chan<- bool) *Snake {
	snake := &Snake{point}

	go snake.move(printListener)

	return snake
}

var _ Printable = &Snake{}
