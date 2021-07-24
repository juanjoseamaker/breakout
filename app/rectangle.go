package app

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Rect struct {
	Color color.Color
	Rect  pixel.Rect
	Vel   pixel.Vec
}

func (r *Rect) Draw(imd *imdraw.IMDraw) {
	imd.Color = r.Color
	imd.Push(r.Rect.Min, r.Rect.Max)
	imd.Rectangle(0)
}

func (r *Rect) Update(dt float64) {
	r.Rect = r.Rect.Moved(r.Vel.Scaled(dt))
}

func (r *Rect) Input(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyRight) {
		r.Vel.X = 100
	}
	if win.Pressed(pixelgl.KeyLeft) {
		r.Vel.X = -100
	}
}
