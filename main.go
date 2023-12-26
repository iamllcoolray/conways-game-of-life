package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iamllcoolray/conways-game-of-life/game"
)

func main() {
	gameOfLife, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(game.SCREENWIDTH, game.SCREENHEIGHT)
	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(gameOfLife); err != nil {
		log.Fatal(err)
	}
}
