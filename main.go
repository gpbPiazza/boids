package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	setEmptyFlocksMapsPosition()
	NewFlock()
	setEachBoidPositionIntoFlockMap()
	s := NewSimulation()

	ebiten.SetWindowSize(s.ScreenWidthPx, s.ScreenWidthPx)
	ebiten.SetWindowTitle("Boids simulation")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
