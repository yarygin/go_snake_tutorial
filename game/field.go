package game

import (
	"gameSnake/cells"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Field struct {
	Width     int
	Height    int
	CellSize  int
	CellSpace int
	Cells     []int
}

func NewField(
	width int,
	height int,
	cellSize int,
	cellSpace int,
) (*Field, error) {
	f := &Field{
		Width:     width,
		Height:    height,
		CellSize:  cellSize,
		CellSpace: cellSpace,
		Cells:     make([]int, width*height),
	}

	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			if isBorder(x, y, f.Width, f.Height) {
				f.setCell(x, y, cells.Wall)
			}
		}
	}

	return f, nil
}

func (f *Field) Update() error {
	return nil
}

func (f *Field) Draw(screen *ebiten.Image) {
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			cell := f.getCell(x, y)
			cellTotalSize := f.CellSize + f.CellSpace
			halfSpace := f.CellSpace / 2
			nx := x*cellTotalSize + halfSpace
			ny := y*cellTotalSize + halfSpace
			ebitenutil.DrawRect(
				screen,
				float64(nx), float64(ny),
				float64(f.CellSize), float64(f.CellSize),
				cells.Palette[cell],
			)
		}
	}
}

func isBorder(x, y, fieldW, fieldH int) bool {
	return x == 0 || x == fieldW-1 || y == 0 || y == fieldH-1
}

func outbound(x, y, w, h int) bool {
	return x < 0 || y < 0 || x >= w || y >= h
}

func (f *Field) getCell(x, y int) int {
	if outbound(x, y, f.Width, f.Height) {
		return 0
	}
	return f.Cells[x*f.Height+y]
}

func (f *Field) setCell(x, y, cell int) {
	if outbound(x, y, f.Width, f.Height) {
		return
	}
	f.Cells[x*f.Height+y] = cell
}
