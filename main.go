package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"github.com/juanjoseamaker/breakout/app"
	"time"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Breakout",
		Bounds: pixel.R(0, 0, 600, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	last := time.Now()

	player := &app.Rect {
		Color: pixel.RGB(255, 0, 0),
		Rect:  pixel.R(0, 0, 100, 50),
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

		imd.Draw(win)
	
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
