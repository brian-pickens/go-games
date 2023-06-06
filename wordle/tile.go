package wordle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	TILE_WIDTH = 20
	TILE_HEIGHT = 24
	MARGIN = 4
)

var TILE_BACKGROUND = color.RGBA{R: 217, G: 139, B: 76, A: 255}
var TILE_FOREGROUND = color.RGBA{R: 76, G: 102, B: 217, A: 255}

type tile struct {
	game	*game
	image	*ebiten.Image
	letter 	rune
	row 	int
	column 	int
}

func newTile(game *game, letter rune, row int, column int) *tile {
	return &tile{
		game: game,
		image: ebiten.NewImage(TILE_WIDTH, TILE_HEIGHT),
		letter: letter,
		row: row,
		column: column,
	}
}

func (t *tile) Draw(boardImage *ebiten.Image) {
	t.image.Fill(TILE_BACKGROUND)
	op := ebiten.DrawImageOptions{}

	op.GeoM.Translate(
		float64((TILE_WIDTH*(t.column-1))+(MARGIN*t.column)),
		float64((TILE_HEIGHT*(t.row-1))+(MARGIN*t.row)))

	if t.letter != 0 {
		text.Draw(
			t.image,
			string(t.letter),
			*t.game.font,
			2,
			(t.image.Bounds().Max.Y)-3,
			TILE_FOREGROUND)
	}

	boardImage.DrawImage(t.image, &op)
}
