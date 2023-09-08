package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidthPx, screenHeightPx = 640, 360
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50}
)

type Simulation struct {
	ScreenWidthPx  int
	ScreenHeightPx int
	Flock          []*Boid
}

func NewSimulation(flock []*Boid) *Simulation {
	return &Simulation{
		ScreenWidthPx:  screenWidthPx,
		ScreenHeightPx: screenHeightPx,
		Flock:          flock,
	}
}

func (s *Simulation) Update() error {
	return nil
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	for _, boid := range s.Flock {
		screen.Set(int(boid.position.x+1), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
	}

}

func (s *Simulation) Layout(_, _ int) (int, int) {
	return s.ScreenWidthPx, s.ScreenHeightPx
}
