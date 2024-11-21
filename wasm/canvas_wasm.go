//go:build js && wasm

package wasm

import (
	"math"
	"syscall/js"

	"github.com/niconielsen24/wasm_boids/boids"
)

func CreateCanvas(w, h int) {
	doc := js.Global().Get("document")

	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", w)
	canvas.Set("height", h)
	canvas.Set("id", "canvas")
  canvas.Set("className","rounded-lg border-2 bg-cyan-200")
	doc.Get("body").Call("appendChild", canvas)
}

func DrawCanvas(boids []boids.Boid, w, h int) {
	doc := js.Global().Get("document")

	canvas := doc.Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")
	for i := range boids {
		drawBoid(boids[i], ctx)
	}
}

func drawBoid(b boids.Boid, ctx js.Value) {
	sx := b.Position.X
	sy := b.Position.Y

	boidSize := 10.0
	dx := b.DirVec.X
	dy := b.DirVec.Y
	angle := math.Atan2(dy, dx)

	frontX := sx + math.Cos(angle)*boidSize // Tip of the boid
	frontY := sy + math.Sin(angle)*boidSize
	leftX := sx + math.Cos(angle+math.Pi*5/6)*boidSize // Left corner
	leftY := sy + math.Sin(angle+math.Pi*5/6)*boidSize
	rightX := sx + math.Cos(angle-math.Pi*5/6)*boidSize // Right corner
	rightY := sy + math.Sin(angle-math.Pi*5/6)*boidSize

	ctx.Call("beginPath")
	ctx.Call("moveTo", frontX, frontY)
	ctx.Call("lineTo", leftX, leftY)
	ctx.Call("lineTo", rightX, rightY)
	ctx.Call("closePath")

	ctx.Set("fillStyle", "blue")
	ctx.Call("fill")
	ctx.Call("stroke")
}

func ClearCanvas(w, h int) {
	doc := js.Global().Get("document")

	canvas := doc.Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")

	ctx.Call("clearRect", 0, 0, w, h)
}
