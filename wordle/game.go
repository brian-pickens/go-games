package wordle

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strings"

	"github.com/brian-pickens/go-games/helpers"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed wordle.txt
var DICTIONARY string
//go:embed tiles_spritesheet.png
var SPRITE_BYTES []byte

const (
	TITLE   string = "Wordle"
	ROWS    int    = 6
	COLUMNS int    = 5
)

type game struct {
	screenWidth   int
	screenHeight  int
	sprite        *ebiten.Image
	font 		  *font.Face
	answer        *answer
	input         [6][5]rune
	currentRow    int
	currentColumn int
	result        string
	guess         bool
	guessState    [COLUMNS]state
}

func StartGame() error {
	game := &game{
		screenWidth:   320,
		screenHeight:  240,
		currentRow:    1,
		currentColumn: 1,
	}

	// Load game sprite
	img, _, err := image.Decode(bytes.NewBuffer(SPRITE_BYTES))
	if err != nil {
		return err
	}
	game.sprite = ebiten.NewImageFromImage(img)

	// Select word from the dictionary
	dict := strings.Split(DICTIONARY, "\n")
	selectedWord := dict[rand.Intn(len(dict))]
	log.Println(selectedWord)
	game.answer = newAnswer(selectedWord)

	// Setup font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	game.font = &font

	// Run Game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for row := 1; row <= ROWS; row++ {
		for column := 1; column <= COLUMNS; column++ {
			tile := newTile(g, g.input[row-1][column-1], row, column)
			tile.Draw(screen)
		}
	}
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
		g.guess, g.guessState = g.answer.guess(g.input[g.currentRow-1])
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
