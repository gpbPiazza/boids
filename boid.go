package main

import (
	"math/rand"
	"time"
)

const (
	boidsCount         = 500
	boidViewRadius     = 13
	adjustVelocityRate = 0.015
)

var (
	flocksMapPositions [screenWidthPx + 1][screenHeightPx + 1]int
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

func (b *Boid) calcAcceleration() Vector2D {
	accel := Vector2D{x: 0, y: 0}
	return accel
}

func (b *Boid) move() {
	// the limit methiod its to ensure the boid will not run faster than 1 px per cycle
	b.velocity = b.velocity.Add(b.calcAcceleration()).LimitVal(-1, 1)
	//set the current position to -1, empty space
	flocksMapPositions[int(b.position.x)][int(b.position.y)] = -1
	// move
	b.position = b.position.Add(b.velocity)
	// fill the new position into the map
	flocksMapPositions[int(b.position.x)][int(b.position.y)] = b.id

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
	for i := 0; i < boidsCount; i++ {
		flock = append(flock, NewBoid(i))
	}
	return flock
}

func setEmptyFlocksMapsPosition() {
	for i, row := range flocksMapPositions {
		for j := range row {
			flocksMapPositions[i][j] = -1
		}
	}
}

func setEachBoidPositionIntoFlockMap(flock []*Boid) {
	for _, boid := range flock {
		flocksMapPositions[int(boid.position.x)][int(boid.position.y)] = boid.id
	}
}
