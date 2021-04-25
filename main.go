package main

import (
	"fmt"
	"gameSnake/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame(800, 600, 17, 17, 16, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ebiten.RunGame(g)
	if err != nil {
		fmt.Println(err)
	}
}
