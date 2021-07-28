package app

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
)

const (
	angleChangeIndex = 10
)

type Ball struct {
	Color  color.Color
	Circle pixel.Circle
	Vel    pixel.Vec
}

func (b *Ball) Draw(imd *imdraw.IMDraw) {
	imd.Color = b.Color
	imd.Push(b.Circle.Center)
	imd.Circle(b.Circle.Radius, 0)
}

func (b *Ball) Update(dt float64) {
	b.Circle.Center = b.Circle.Center.Add(b.Vel.Scaled(dt))
}

func (b *Ball) Bounce(dir pixel.Vec) {
	b.Vel = b.Vel.ScaledXY(dir)
}

func (b *Ball) IsCollidingBounds(bounds pixel.Vec) {
	if b.Circle.Center.X - b.Circle.Radius < 0 || b.Circle.Center.X + b.Circle.Radius > bounds.X {
		b.Bounce(pixel.V(-1, 1))
	}

	if b.Circle.Center.Y + b.Circle.Radius > bounds.Y {
		b.Bounce(pixel.V(1, -1))
	}
}

func (b *Ball) IsCollidingRect(rect pixel.Rect) bool {
	colliding := b.Circle.IntersectRect(rect)

	if !colliding.Eq(pixel.ZV) {
		if colliding.X != 0 {
			b.Vel.X *= -1
		} else {
			b.Vel.Y *= -1
		}
		
		return true
	}
	
	return false
}

func (b *Ball) IsCollidingPlataform(rect pixel.Rect, vel pixel.Vec) {
	colliding := b.Circle.IntersectRect(rect)
	
	if !colliding.Eq(pixel.ZV) && b.Vel.Y < 0 {
		b.Vel.Y *= -1
	}
}
