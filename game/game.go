package game

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	SCREENWIDTH  = 720
	SCREENHEIGHT = 480

	SCALE = 2
)

var (
	pink color.RGBA = color.RGBA{255, 0, 255, 255}

	grid   [SCREENWIDTH][SCREENHEIGHT]uint8 = [SCREENWIDTH][SCREENHEIGHT]uint8{}
	buffer [SCREENWIDTH][SCREENHEIGHT]uint8 = [SCREENWIDTH][SCREENHEIGHT]uint8{}

	timer uint16 = 0
)

func NewGame() (*Game, error) {
	for row := 0; row < SCREENWIDTH; row++ {
		for col := 30; col < SCREENHEIGHT; col++ {
			if rand.Float32() < 0.5 {
				grid[row][col] = 1
			}
		}
	}

	game := &Game{}
	var err error
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *Game) Update() error {
	timer++

	for row := 1; row < SCREENWIDTH-1; row++ {
		for col := 1; col < SCREENHEIGHT-1; col++ {
			buffer[row][col] = 0
			neighbors := grid[row-1][col-1] + grid[row-1][col+0] + grid[row-1][col+1] + grid[row+0][col-1] + grid[row+0][col+1] + grid[row+1][col-1] + grid[row+1][col+0] + grid[row+1][col+1]

			if grid[row][col] == 0 && neighbors == 3 {
				buffer[row][col] = 1
			} else if neighbors > 3 || neighbors < 2 {
				buffer[row][col] = 0
			} else {
				buffer[row][col] = grid[row][col]
			}
		}
	}

	temp := buffer
	buffer = grid
	grid = temp

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := 0; row < SCREENWIDTH; row++ {
		for col := 30; col < SCREENHEIGHT; col++ {
			if grid[row][col] > 0 {
				screen.Set(row, col, pink)
			}
		}
	}

	timerStr := fmt.Sprintf("Timer: %d", timer)
	ebitenutil.DebugPrint(screen, timerStr)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
