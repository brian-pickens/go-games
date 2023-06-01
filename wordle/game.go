package wordle

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TITLE   string = "Wordle"
	ROWS    int    = 6
	COLUMNS int    = 5
)

type game struct{
	screenWidth int
	screenHeight int
	input [6][5][]rune
	currentRow int
	currentColumn int
	result string
}

func StartGame() (error) {
	game := game{
		screenWidth: 320,
		screenHeight: 240,
		currentRow: 1,
		currentColumn: 1,
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	var text string
	for row := 1; row <= ROWS; row++ {
		for column := 1; column <= COLUMNS; column++ {
			text += string(g.input[row-1][column-1])
		}
		text += "\n"
	}
	text += g.result
	ebitenutil.DebugPrint(screen, text)
}

func (g *game) Update() error {
	var input []rune
	input = ebiten.AppendInputChars(input[:0])

	if g.currentRow > ROWS {
		g.result = "END"
		return nil
	}
	
	if (len(input) > 0) {
		g.input[g.currentRow-1][g.currentColumn-1] = input
		g.currentColumn++
	}

	if (g.currentColumn > COLUMNS) {
		g.currentColumn = 1
		g.currentRow++
	}

	return nil
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.screenWidth, g.screenHeight
}