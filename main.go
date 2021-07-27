package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
	"github.com/juanjoseamaker/breakout/app"
	"time"
)

const (
	Gaps = 16
	BrickSizeX = 90
	BrickSizeY = 50
	WindowSizeX = 650
	WindowSizeY = 650
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Breakout",
		Bounds: pixel.R(0, 0, WindowSizeX, WindowSizeY),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	last := time.Now()

	level := make([][]uint8, 10)
	for i := range level {
		level[i] = make([]uint8, 10)
	}
	
	app.LoadLevel(level, "./level01.csv")

	player := &app.Rect {
		Color: pixel.RGB(255, 0, 0),
		Rect:  pixel.R(0, 0, 150, 25),
		Vel:   pixel.V(0, 0),
	}

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
	
		win.Clear(colornames.Gray)
		imd.Clear()

		player.Input(win)
		player.Update(dt)
		player.Draw(imd)

		for i := range level {
			for j := range level[i] {
				var color color.Color
				
				switch level[i][j] {
				case 0:
					continue
				case 1:
					color = pixel.RGB(255, 0, 0)
				case 2:
					color = pixel.RGB(0, 255, 0)
				case 3:
					color = pixel.RGB(0, 0, 255)
				}
				
				app.DrawRect(
					imd,
					pixel.R(
						float64(Gaps*(j+1) + BrickSizeX*j),
						WindowSizeY	- float64(Gaps*(i+1) + BrickSizeY*(i+1)),
						float64(Gaps*(j+1) + BrickSizeX*(j+1)),
						WindowSizeY	- float64(Gaps*(i+1) + BrickSizeY*i)),
					color,
				)
			}
		}

		imd.Draw(win)
	
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
