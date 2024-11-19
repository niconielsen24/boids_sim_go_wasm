package main

import (
	"fmt"

	//"github.com/niconielsen24/wasm_boids/boids"
	"github.com/niconielsen24/wasm_boids/quad_tree"
)

func main() {
	p1 := quadtree.Point{X: 0, Y: 0}
	p2 := quadtree.Point{X: 8, Y: 8}
	t := quadtree.QuadInit(p1, p2)
	a := quadtree.Node{
		Position: quadtree.Point{X: 1, Y: 1},
		Data:     1,
	}
	b := quadtree.Node{
		Position: quadtree.Point{X: 2, Y: 5},
		Data:     2,
	}
	c := quadtree.Node{
		Position: quadtree.Point{X: 7, Y: 6},
		Data:     3,
	}
	t.Insert(&a)
	t.Insert(&b)
	t.Insert(&c)
  fmt.Printf("Node a : %v\n", t.Search(quadtree.Point{X: 1, Y: 1}).Data)
  fmt.Printf("Node b : %v\n", t.Search(quadtree.Point{X: 2, Y: 5}).Data)
  fmt.Printf("Node c : %v\n", t.Search(quadtree.Point{X: 7, Y: 6}).Data)
  fmt.Printf("Non existing node : %v", t.Search(quadtree.Point{X: 5, Y: 5}))
}
