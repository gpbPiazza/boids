package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	boidsCount         = 500
	boidViewRadius     = 13
	adjustVelocityRate = 0.015
)

var (
	flockMapPositions = [screenWidth + 1][screenHeight + 1]int{}
	flock             = make(map[int]*Boid, 0)
	rwLocker          = sync.RWMutex{}
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func NewBoid(id int) *Boid {
	b := &Boid{
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
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
	upperView := b.position.AddVal(boidViewRadius)
	lowerView := b.position.AddVal(-boidViewRadius)
	// all variables with prefix all here mean all elements inside of viewBox, inside of Boid View Radius
	allBoidsVelocity := Vector2D{x: 0, y: 0}
	allBoidsPosition := Vector2D{x: 0, y: 0}
	allBoidsSeparation := Vector2D{x: 0, y: 0}
	accel := Vector2D{x: 0, y: 0}
	boidsCount := 0.0

	// Começa na posição mais baixa de X e anda até o máximo ou limite da tela
	// iteração entre todos os elementos dentros da viewBox do Boid

	rwLocker.RLock()
	for i := math.Max(lowerView.x, 0); i <= math.Min(upperView.x, screenWidth); i++ {
		for j := math.Max(lowerView.y, 0); j <= math.Min(upperView.y, screenHeight); j++ {
			otherBoidId := flockMapPositions[int(i)][int(j)]
			if otherBoidId != -1 && b.id != otherBoidId {
				otherBoid := flock[otherBoidId]
				dist := otherBoid.position.Distance(b.position)
				if dist < boidViewRadius {
					boidsCount++
					allBoidsVelocity = allBoidsVelocity.Add(otherBoid.velocity)
					allBoidsPosition = allBoidsPosition.Add(otherBoid.position)
					separation := b.position.Subtract(otherBoid.position).DivisionVal(dist)
					allBoidsSeparation = allBoidsSeparation.Add(separation)
				}
			}
		}
	}
	rwLocker.RUnlock()

	if boidsCount > 0 {
		avgVelocity := allBoidsVelocity.DivisionVal(boidsCount)
		avgPosition := allBoidsPosition.DivisionVal(boidsCount)
		accelAligment := avgVelocity.Subtract(b.velocity).MultiplyVal(adjustVelocityRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyVal(adjustVelocityRate)
		accelSepartion := allBoidsSeparation.MultiplyVal(adjustVelocityRate)
		accel = accel.Add(accelAligment).Add(accelCohesion).Add(accelSepartion)
	}

	return accel
}

func (b *Boid) move() {
	accel := b.calcAcceleration()

	rwLocker.Lock()
	// the limit methiod its to ensure the boid will not run faster than 1 px per cycle
	b.velocity = b.velocity.Add(accel).LimitVal(-1, 1)
	//set the current position to -1, empty space
	flockMapPositions[int(b.position.x)][int(b.position.y)] = -1
	// move
	b.position = b.position.Add(b.velocity)
	// fill the new position into the map
	flockMapPositions[int(b.position.x)][int(b.position.y)] = b.id

	nextPosition := b.position.Add(b.velocity)
	// Ensure the boid will not pass throught the screen limits
	// If reaches the limit will change the direction
	if nextPosition.x >= screenWidth || nextPosition.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}

	if nextPosition.y >= screenHeight || nextPosition.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}
	rwLocker.Unlock()
}

func NewFlock() {
	for i := 0; i < boidsCount; i++ {
		boid := NewBoid(i)
		flock[boid.id] = boid
	}
}

func setEmptyFlocksMapsPosition() {
	for i, row := range flockMapPositions {
		for j := range row {
			flockMapPositions[i][j] = -1
		}
	}
}

func setEachBoidPositionIntoFlockMap() {
	for _, boid := range flock {
		flockMapPositions[int(boid.position.x)][int(boid.position.y)] = boid.id
	}
}
