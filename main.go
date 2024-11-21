//go:build js && wasm

package main

import (
	"github.com/niconielsen24/wasm_boids/boids"
	"github.com/niconielsen24/wasm_boids/wasm"
	"math/rand"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	numBoids := 100
	//viewAngle := 120.0 // Field of view in degrees
	//viewDistance := 50.0
	width := 800
	height := 600

	boidFlock := make([]boids.Boid, numBoids)
	for i := 0; i < numBoids; i++ {
		boidFlock[i] = boids.Boid{
			Position: boids.Position{
				X: randFloat(0, float64(width)),
				Y: randFloat(0, float64(height)),
			},
			DirVec: boids.Vector{
				X: randFloat(-1, 1),
				Y: randFloat(-1, 1),
			},
		}
		boidFlock[i].DirVec.Normalize()
	}

	wasm.DrawCanvas(boidFlock, width, height)
}
