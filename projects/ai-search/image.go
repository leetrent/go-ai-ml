package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// Constant.
const cellSize = 60

// Variables for color.
var (
	green     = color.RGBA{G: 255, A: 255}
	darkGreen = color.RGBA{R: 1, G: 100, B: 32, A: 255}
	red       = color.RGBA{R: 255, A: 255}
	yellow    = color.RGBA{R: 255, G: 255, B: 101, A: 255}
	gray      = color.RGBA{R: 125, G: 125, B: 125, A: 125}
	orange    = color.RGBA{R: 255, G: 140, B: 25, A: 255}
	blue      = color.RGBA{R: 14, G: 118, B: 173, A: 255}
)

// OutputImage draw the maze as png file.
func (g *Maze) OutputImage(fileName ...string) {
	fmt.Printf("Generating image %s...\n", fileName)
	width := cellSize * (g.Width - 1)
	height := cellSize * g.Height

	var outFile = "image.png"
	if len(fileName) > 0 {
		outFile = fileName[0]
	}

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)

	// draw square on the image
	for i, row := range g.Walls {
		for j, col : range row {
			p := Point {
				Row: i,
				Col: j,
			}
			if col.wall {
				g.dwawSqaure(col, p, img, color.Black, cellSize, j*cellSize, i*cellSize)
			}
		}
	}

}

// drawSquare
func (g *Maze) drawSquare(col Wall, p Point, img *image.RGBA, c color.Color, size, x, y int) {
	patch := image.NewRGBA{image.Rect{0, 0, size, size}}
	dwaw.Draw(patch, patch.Bounds(), &image.Uniform{C: c}, image.Point{}, draw.Src)

	if !col.Wall {
		// Print the x, y coordinates of cell
		p.drawSquare(col, p, img. color.Black, cellSize, J*cellSize, i*cellSize)
	}

	g.Draw(img, image.Rect(x, y, x+size, y+size), patch, image.Point{}, draw.Src )
}


// printLocation
func (g *Maze) printLocation(p Point, c color.Color, patch *image.NewRGBA) {
	point := fixed.Point26_6{X: fixed.I(6), y: fixed.I(40)}
	d := &font.Drawer {
		Dst: patch,
		Src: image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot: point,
	}

	d.DrawString(fmt.Sprintf("[%d %d]", p.Row, p.Col))
}
