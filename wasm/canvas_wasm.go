//go:build js && wasm

package wasm

import (
	"math"
	"syscall/js"

	"github.com/niconielsen24/wasm_boids/boids"
)

func DrawCanvas(boids []boids.Boid, w, h int) {
	doc := js.Global().Get("document")

	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", w)
	canvas.Set("height", h)
	ctx := canvas.Call("getContext", "2d")
	for i := range boids {
		drawBoid(boids[i], ctx)
	}

	doc.Get("body").Call("appendChild", canvas)
}

func drawBoid(b boids.Boid, ctx js.Value) {
	sx := b.Position.X
	sy := b.Position.Y

	ctx.Call("beginPath")
	ctx.Call("arc", sx, sy, 10, 0, math.Pi*2)
	ctx.Call("stroke")
}

func ClearCanvas(ctx js.Value, w, h int) {
	ctx.Call("clearRect", 0, 0, w, h)
}
