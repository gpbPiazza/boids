package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func NewBoid(id int) *Boid {
	b := &Boid{
		position: Vector2D{x: rand.Float64() * screenWidthPx, y: rand.Float64() * screenHeightPx},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       id,
	}

	go b.fly()

	return b
}

func (b *Boid) fly() {
	for {
		b.move()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) move() {
	b.position = b.position.Add(b.velocity)

	nextPosition := b.position.Add(b.velocity)
	// Ensure the boid will not pass throught the screen limits
	// If reaches the limit will change the direction
	if nextPosition.x >= screenWidthPx || nextPosition.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}

	if nextPosition.y >= screenHeightPx || nextPosition.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}
}

func NewFlock() []*Boid {
	var flock []*Boid
	for i := 0; i < 500; i++ {
		flock = append(flock, NewBoid(i))
	}
	return flock
}
