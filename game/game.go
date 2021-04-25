package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Field *Field
	Snake *Snake
}

func NewGame(
	sizeW int,
	sizeH int,
	fieldW int,
	fieldH int,
	cellSize int,
	cellSpace int,
) (*Game, error) {
	g := &Game{}

	ebiten.SetWindowSize(sizeW, sizeH)
	ebiten.SetWindowResizable(true)

	field, err := NewField(fieldW, fieldH, cellSize, cellSpace)
	if err != nil {
		return nil, err
	}

	g.Field = field
	g.Snake = NewSnake(field)

	return g, nil
}

func (g *Game) Update() error {
	if err := g.Field.Update(); err != nil {
		return err
	}
	if err := g.Snake.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Field.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	cellTotalSize := g.Field.CellSize + g.Field.CellSpace
	return g.Field.Width * cellTotalSize, g.Field.Height * cellTotalSize
}
