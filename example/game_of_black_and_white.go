package main

import (
	"image"
	"image/color"

	"github.com/DumDumGeniuss/ggol"
)

type gameOfBlackAndWhiteArea struct {
	HasLiveCell bool
}

var initialGameOfBlackAndWhiteArea gameOfBlackAndWhiteArea = gameOfBlackAndWhiteArea{
	HasLiveCell: false,
}

func gameOfBlackAndWhiteAreaIterator(
	coord *ggol.Coordinate,
	area *gameOfBlackAndWhiteArea,
	getAdjacentArea ggol.AdjacentAreaGetter[gameOfBlackAndWhiteArea],
) (nextArea *gameOfBlackAndWhiteArea) {
	newArea := *area

	if newArea.HasLiveCell {
		newArea.HasLiveCell = false
		return &newArea
	} else {
		newArea.HasLiveCell = true
		return &newArea
	}
}

func initializeGameOfBlackAndWhiteField(g ggol.Game[gameOfBlackAndWhiteArea]) {
	size := g.GetSize()
	for x := 0; x < size.Width; x++ {
		for y := 0; y < size.Height; y++ {
			c := ggol.Coordinate{X: x, Y: y}
			g.SetArea(&c, &gameOfBlackAndWhiteArea{HasLiveCell: (x+y)%3 == 0})
		}
	}
}

func drawGameOfBlackAndWhiteArea(coord *ggol.Coordinate, area *gameOfBlackAndWhiteArea, unit int, image *image.Paletted, palette *[]color.Color) {
	if !area.HasLiveCell {
		return
	}
	for i := 0; i < unit; i += 1 {
		for j := 0; j < unit; j += 1 {
			image.Set(coord.X*unit+i, coord.Y*unit+j, (*palette)[1])
		}
	}
}

func executeGameOfBlackAndWhite() {
	size := ggol.Size{Width: 50, Height: 50}
	game, _ := ggol.New(&size, &initialGameOfBlackAndWhiteArea)
	game.SetAreaIterator(gameOfBlackAndWhiteAreaIterator)
	initializeGameOfBlackAndWhiteField(game)

	var gameOfBlackAndWhitePalette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var images []*image.Paletted
	var delays []int
	unit := 10
	iterationsCount := 100
	duration := 100

	for i := 0; i < iterationsCount; i += 1 {
		newImage := image.NewPaletted(image.Rect(0, 0, size.Width*unit, size.Height*unit), gameOfBlackAndWhitePalette)
		for x := 0; x < size.Width; x += 1 {
			for y := 0; y < size.Height; y += 1 {
				coord := &ggol.Coordinate{X: x, Y: y}
				area, _ := game.GetArea(coord)
				drawGameOfBlackAndWhiteArea(coord, area, unit, newImage, &gameOfBlackAndWhitePalette)
			}
		}
		images = append(images, newImage)
		delays = append(delays, duration)
		game.Iterate()
	}

	outputGif("output/game_of_black_and_white.gif", images, delays)
}
