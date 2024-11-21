//go:build js && wasm

package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/niconielsen24/wasm_boids/boids"
	quadtree "github.com/niconielsen24/wasm_boids/quad_tree"
	"github.com/niconielsen24/wasm_boids/wasm"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	numBoids := 400
	viewAngle := 120.0 // Field of view in degrees
	viewDistance := 20.0
	width := 600
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

	hd := float64(width / 2)
	b := quadtree.Boundary{
		Center: quadtree.Point{
			X: hd,
			Y: hd,
		},
		HalfDim: hd,
	}
	qt := quadtree.InitQuad[boids.Boid](&b)
	wasm.CreateCanvas(width, height)

	go func() {
		for {
			// Initialize quadtree
			qt.Clear()
			// Insert boids into quadtree
			for i := range boidFlock {
				qt.Insert(&quadtree.UserPoint[boids.Boid]{
					X:        boidFlock[i].Position.X,
					Y:        boidFlock[i].Position.Y,
					UserData: boidFlock[i],
				})
			}

			wasm.ClearCanvas(width, height)
			wasm.DrawCanvas(boidFlock, width, height)

			var mutex sync.Mutex
			chunkSize := (len(boidFlock) + 3) / 4
			var wg sync.WaitGroup

			for i := 0; i < 4; i++ {
				start := i * chunkSize
				end := start + chunkSize
				if end > len(boidFlock) {
					end = len(boidFlock)
				}

				wg.Add(1)
				go func(start, end int) {
					defer wg.Done()
					localUpdates := make([]boids.Boid, end-start)
					for j := start; j < end; j++ {
						b := &boidFlock[j]
						r := quadtree.Boundary{
							Center:  quadtree.Point{X: b.Position.X, Y: b.Position.Y},
							HalfDim: viewDistance * 2,
						}

						query := qt.Query(&r)
						var neighbors []boids.Boid
						for _, point := range query {
							neighbors = append(neighbors, point.UserData)
						}

						b.Update(neighbors, viewAngle, viewDistance)
						b.KeepInBounds(float64(width), float64(height))
						localUpdates[j-start] = *b
					}

					mutex.Lock()
					copy(boidFlock[start:end], localUpdates)
					mutex.Unlock()
				}(start, end)
			}

			wg.Wait()
			time.Sleep(16 * time.Millisecond)
		}
	}()

	select {}
}
