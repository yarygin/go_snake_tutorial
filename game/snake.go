package game

import (
	"gameSnake/cells"
	"gameSnake/util"
)

type Snake struct {
	Field     *Field
	Body      []util.Index
	Direction int
}

const (
	DirectionUp = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

func NewSnake(field *Field) *Snake {
	cx, cy := field.Width/2, field.Height/2
	s := &Snake{
		Field: field,
		Body: []util.Index{
			{cx - 1, cy},
			{cx, cy},
			{cx + 1, cy},
		},
		Direction: DirectionUp,
	}
	for _, idx := range s.Body {
		s.Field.setCell(idx.I, idx.J, cells.Snake)
	}
	return s
}

func (s *Snake) Update() error {
	return nil
}
