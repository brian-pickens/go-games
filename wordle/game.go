package wordle

import (
	_ "embed"
	"log"
	"math/rand"
	"strings"

	"github.com/brian-pickens/go-games/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	input [6][5]rune
	currentRow int
	currentColumn int
	result string
	guess bool
	guessState [COLUMNS]state
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
	log.Println(selectedWord)
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

	// Undo typed characters
	if g.currentColumn > 1 &&
	   inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		g.input[g.currentRow-1][g.currentColumn-2] = rune(0)
		g.currentColumn--
	}

	// Guess the answer when the Enter key is pressed
	if g.currentColumn > COLUMNS &&
	   ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.guess, g.guessState = g.answer.Guess(g.input[g.currentRow-1])
		log.Println(g.guess, g.guessState)
		g.currentColumn = 1
		g.currentRow++
	}

	// Hold if all letters for the current row are guessed
	if g.currentColumn > COLUMNS {
		return nil
	}

	// Win condition
	if g.guess {
		g.result = "WIN"
		return nil
	}

	// Loose condition
	if g.currentRow > ROWS {
		g.result = "LOOSE"
		return nil
	}

	// Handle User Input
	input = ebiten.AppendInputChars(input[:0])
	if (len(input) > 0) {
		columns := helpers.Min(len(input), (COLUMNS-g.currentColumn+1))
		for i := 0; i < columns; i++ {
			g.input[g.currentRow-1][g.currentColumn-1+i] = input[i]
		}
		g.currentColumn += columns
	}

	return nil
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

