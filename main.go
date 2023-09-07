package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	flock := NewFlock()
	s := NewSimulation(flock)

	ebiten.SetWindowSize(s.ScreenWidthPx, s.ScreenWidthPx)
	ebiten.SetWindowTitle("Boids simulation")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
