package quadtree

type Point struct {
	X int64
	Y int64
}

type Node struct {
	Position Point
	Data     int // boid??
}

type Quad struct {
	TopLeft     Point
	BottomRight Point
	Node        *Node

	TopLeftTree  *Quad
	TopRightTree *Quad
	BotLeftTree  *Quad
	BotRightTree *Quad
}

// Initialize quad tree, where tl and br are points
// that define the  bounds of this tree
func QuadInit(tl, br Point) *Quad {
	return &Quad{
		TopLeft:     tl,
		BottomRight: br,

		Node:         nil,
		TopLeftTree:  nil,
		TopRightTree: nil,
		BotLeftTree:  nil,
		BotRightTree: nil,
	}
}

func (qt *Quad) inBoundary(p Point) bool {
	return (p.X >= qt.TopLeft.X &&
		p.X <= qt.BottomRight.X &&
		p.Y >= qt.TopLeft.Y &&
		p.Y <= qt.BottomRight.Y)
}

func abs(n int64) int64 {
	if n < 0 {
		return n * -1
	}
	return n
}

func (qt *Quad) Insert(n *Node) {
	if n == nil {
		return
	}

	// current quad cannot contain it
	if !qt.inBoundary(n.Position) {
		return
	}

	// unit area quad cannot be subdivided further
	if abs(qt.TopLeft.X-qt.BottomRight.X) <= 1 &&
		abs(qt.TopLeft.Y-qt.BottomRight.Y) <= 1 {
		if qt.Node == nil {
			qt.Node = n
      return
		}
	}

	if (qt.TopLeft.X+qt.BottomRight.X)/2 >= n.Position.X {
		if (qt.TopLeft.Y+qt.BottomRight.Y)/2 >= n.Position.Y {
			// We are on TopLeftTree area
			if qt.TopLeftTree == nil {
				qt.TopLeftTree = QuadInit(
					Point{
						X: qt.TopLeft.X,
						Y: qt.TopLeft.Y,
					},
					Point{
						X: (qt.TopLeft.X + qt.BottomRight.X) / 2,
						Y: (qt.TopLeft.Y + qt.BottomRight.Y) / 2,
					},
				)
			}
			qt.TopLeftTree.Insert(n)
		} else {
			// We are on BotLeftTree area
			if qt.BotLeftTree == nil {
				qt.BotLeftTree = QuadInit(
					Point{
						X: qt.TopLeft.X,
						Y: (qt.TopLeft.Y + qt.BottomRight.Y) / 2,
					},
					Point{
						X: (qt.TopLeft.X + qt.BottomRight.X) / 2,
						Y: qt.BottomRight.Y,
					},
				)
			}
			qt.BotLeftTree.Insert(n)
		}
	} else {
		if (qt.TopLeft.Y+qt.BottomRight.Y)/2 >= n.Position.Y {
			// We are on TopRightTree area
			if qt.TopRightTree == nil {
				qt.TopRightTree = QuadInit(
					Point{
						X: (qt.TopLeft.X + qt.BottomRight.X) / 2,
						Y: qt.TopLeft.Y,
					},
					Point{
						X: qt.BottomRight.X,
						Y: (qt.TopLeft.Y + qt.BottomRight.Y) / 2,
					},
				)
			}
			qt.TopRightTree.Insert(n)
		} else {
			// We are on BotRightTree area
			if qt.BotRightTree == nil {
				qt.BotRightTree = QuadInit(
					Point{
						X: (qt.TopLeft.X + qt.BottomRight.X) / 2,
						Y: (qt.TopLeft.Y + qt.BottomRight.Y) / 2,
					},
					Point{
						X: qt.BottomRight.X,
						Y: qt.BottomRight.Y,
					},
				)
			}
			qt.BotRightTree.Insert(n)
		}
	}
}

func (qt *Quad) Search(p Point) *Node {
	// Outside current quad bounds
	if !qt.inBoundary(p) {
		return nil
	}

	// Unit lenght quad cannot be subdivided
	if qt.Node != nil {
		return qt.Node
	}

	if (qt.TopLeft.X+qt.BottomRight.X)/2 >= p.X {
		if (qt.TopLeft.Y+qt.BottomRight.Y)/2 >= p.Y {
			// We are on TopLeftTree area
			if qt.TopLeftTree == nil {
				return nil
			}
			return qt.TopLeftTree.Search(p)
		} else {
			// We are on BotLeftTree area
			if qt.BotLeftTree == nil {
				return nil
			}
			return qt.BotLeftTree.Search(p)
		}
	} else {
		if (qt.TopLeft.Y+qt.BottomRight.Y)/2 >= p.Y {
			// We are on TopRightTree area
			if qt.TopRightTree == nil {
				return nil
			}
			return qt.TopRightTree.Search(p)
		} else {
      // We are on BotRightTree area
      if qt.BotRightTree == nil {
        return nil
      }
      return qt.BotRightTree.Search(p)
    }
	}
}
