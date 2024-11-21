package quadtree

type Point struct {
	X float64
	Y float64
}

type UserPoint[T any] struct {
	X        float64
	Y        float64
	UserData T
}

type Boundary struct {
	Center  Point
	HalfDim float64
}

type Quad[T any] struct {
	Bound  Boundary
	Points []UserPoint[T]

	nw *Quad[T]
	ne *Quad[T]
	sw *Quad[T]
	se *Quad[T]
}

func (b *Boundary) containsPoint(p Point) bool {
	return (b.Center.X+b.HalfDim >= p.X && b.Center.X-b.HalfDim <= p.X &&
		b.Center.Y+b.HalfDim >= p.Y && b.Center.Y-b.HalfDim <= p.Y)
}

func (b *Boundary) intersects(other *Boundary) bool {
	return !(b.Center.X+b.HalfDim < other.Center.X-other.HalfDim ||
		b.Center.X-b.HalfDim > other.Center.X+other.HalfDim ||
		b.Center.Y+b.HalfDim < other.Center.Y-other.HalfDim ||
		b.Center.Y-b.HalfDim > other.Center.Y+other.HalfDim)
}

func InitQuad[T any](b *Boundary) *Quad[T] {
	return &Quad[T]{
		Bound:  *b,
		Points: []UserPoint[T]{},
		nw:     nil,
		ne:     nil,
		sw:     nil,
		se:     nil,
	}
}

func (qt *Quad[T]) Subdivide() {
	half := qt.Bound.HalfDim / 2
	qt.nw = InitQuad[T](&Boundary{
		Center:  Point{qt.Bound.Center.X - half, qt.Bound.Center.Y + half},
		HalfDim: half,
	})
	qt.ne = InitQuad[T](&Boundary{
		Center:  Point{qt.Bound.Center.X + half, qt.Bound.Center.Y + half},
		HalfDim: half,
	})
	qt.sw = InitQuad[T](&Boundary{
		Center:  Point{qt.Bound.Center.X - half, qt.Bound.Center.Y - half},
		HalfDim: half,
	})
	qt.se = InitQuad[T](&Boundary{
		Center:  Point{qt.Bound.Center.X + half, qt.Bound.Center.Y - half},
		HalfDim: half,
	})
}

func (qt *Quad[T]) Insert(up *UserPoint[T]) bool {

	if !qt.Bound.containsPoint(Point{X: up.X, Y: up.Y}) {
		return false
	}

	if len(qt.Points) < 4 && qt.nw == nil {
		qt.Points = append(qt.Points, *up)
		return true
	}

	if qt.nw == nil {
		qt.Subdivide()
	}

	if qt.nw.Insert(up) {
		return true
	}
	if qt.ne.Insert(up) {
		return true
	}
	if qt.sw.Insert(up) {
		return true
	}
	if qt.se.Insert(up) {
		return true
	}
	return false
}

func (qt *Quad[T]) Query(b *Boundary) []UserPoint[T] {
  var pointsInRange []UserPoint[T]
  
  if !qt.Bound.intersects(b) {
    return pointsInRange
  }

  for i := range qt.Points {
    p := qt.Points[i]
    if qt.Bound.containsPoint(Point{X: p.X, Y: p.Y}) {
      pointsInRange = append(pointsInRange, p)
    }
  }

  if qt.nw == nil {
    return pointsInRange
  }

  pointsInRange = append(pointsInRange, qt.nw.Query(b)...)
  pointsInRange = append(pointsInRange, qt.ne.Query(b)...)
  pointsInRange = append(pointsInRange, qt.sw.Query(b)...)
  pointsInRange = append(pointsInRange, qt.se.Query(b)...)

  return pointsInRange
}
