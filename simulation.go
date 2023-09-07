package boids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidthPx, screenHeightPx = 640, 360
	boidsCount                    = 500
)

var green = color.RGBA{R: 10, G: 50, B: 255, A: 0}

type Simulation struct {
	ScreenWidthPx  int
	ScreenHeightPx int
}

func NewSimulation() *Simulation {
	return &Simulation{
		ScreenWidthPx:  screenWidthPx,
		ScreenHeightPx: screenHeightPx,
	}
}

func (s *Simulation) Update() error {
	return nil
}

func (s *Simulation) Simulate(screen *ebiten.Image) {

}

func (s *Simulation) Layout() (int, int) {
	return s.ScreenWidthPx, s.ScreenHeightPx
}
