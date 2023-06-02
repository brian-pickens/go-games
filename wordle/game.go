package wordle

import (
	_ "embed"
	"log"
	"math/rand"
	"strings"

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
	answer answer
	input [6][5][]rune
	currentRow int
	currentColumn int
	result string
}

//go:embed wordle.txt
var DICTIONARY string

func StartGame() (error) {
	game := game{
		screenWidth: 320,
		screenHeight: 240,
		currentRow: 1,
		currentColumn: 1,
	}

	dict := strings.Split(DICTIONARY, "\n")
	selectedWord := dict[rand.Intn(len(dict))]
	game.answer = NewAnswer(selectedWord)
	
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

	if g.currentColumn > COLUMNS &&
	   ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.currentColumn = 1
		g.currentRow++
	}
	if g.currentColumn > COLUMNS {
		return nil
	}

	if g.currentRow > ROWS {
		g.result = "END"
		return nil
	}

	input = ebiten.AppendInputChars(input[:0])

	if (len(input) > 0) {
		g.input[g.currentRow-1][g.currentColumn-1] = input
		g.currentColumn++
	}

	return nil
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

