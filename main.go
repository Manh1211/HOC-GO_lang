package main

import (
	"fmt"
	"log"
)

//Point is a 2d point
type Point struct {
	X int
	Y int
}

//Square is a square
type Square struct {
	Center Point
	Length int
}

//Moved moves the point
func (p *Point) Moved(dx int, dy int) {
	p.X += dx
	p.Y += dy

}

//NewSuqare return a new square
func NewSuqare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length must be > 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

//S.method have prototype P.method
func (s *Square) Move(dx int, dy int) {
	s.Center.Moved(dx, dy) // s --> centrer --> Move
}

//Area return the square are
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSuqare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
