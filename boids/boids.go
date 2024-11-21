package boids

import "math"

type Position struct {
	X float64
	Y float64
}

type Vector struct {
	X float64
	Y float64
}

type ForceScalar float64

type Boid struct {
	Position Position
	DirVec   Vector
	DirAngle float64
}

func (v *Vector) Normalize() {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if len != 0 {
		v.X /= len
		v.Y /= len
	}
}

func (v *Vector) add(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}

func (p *Position) add(other *Position) {
	p.X += other.X
	p.Y += other.Y
}

func InViewRange(origin, other *Boid, viewAngle, viewDistance float64) bool {
    delta := Vector{
        X: other.Position.X - origin.Position.X,
        Y: other.Position.Y - origin.Position.Y,
    }

    distance := math.Sqrt(delta.X*delta.X + delta.Y*delta.Y)
    if distance > viewDistance {
        return false
    }

    delta.Normalize()
    origin.DirVec.Normalize()

    dot := delta.X*origin.DirVec.X + delta.Y*origin.DirVec.Y
    angle := math.Acos(dot) * (180 / math.Pi) // Convert to degrees

    return angle <= viewAngle/2
}


func (b *Boid) Update(neighbors []Boid, viewAngle, viewDistance float64) {
    cohesion := Vector{}
    separation := Vector{}
    alignment := Vector{}
    neighborCount := 0

    for _, neighbor := range neighbors {
        if InViewRange(b, &neighbor, viewAngle, viewDistance) {
            neighborCount++

            // Distance-based force scalar
            delta := Vector{
                X: neighbor.Position.X - b.Position.X,
                Y: neighbor.Position.Y - b.Position.Y,
            }
            distance := math.Sqrt(delta.X*delta.X + delta.Y*delta.Y)

            // Cohesion: Move toward the average position of neighbors
            cohesion.X += neighbor.Position.X
            cohesion.Y += neighbor.Position.Y

            // Separation: Avoid too-close neighbors
            if distance > 0 {
                separation.X += (b.Position.X - neighbor.Position.X) / distance 
                separation.Y += (b.Position.Y - neighbor.Position.Y) / distance
            }

            // Alignment: Match the average direction of neighbors
            alignment.X += neighbor.DirVec.X 
            alignment.Y += neighbor.DirVec.Y
        }
    }

    if neighborCount > 0 {
        // Average cohesion force
        cohesion.X /= float64(neighborCount)
        cohesion.Y /= float64(neighborCount)
        cohesion.X -= b.Position.X
        cohesion.Y -= b.Position.Y
        cohesion.Normalize()

        // Normalize separation
        separation.Normalize()

        // Average alignment force
        alignment.X /= float64(neighborCount)
        alignment.Y /= float64(neighborCount)
        alignment.Normalize()
    }

    // Combine forces with weights
    b.DirVec.X += cohesion.X + separation.X + alignment.X
    b.DirVec.Y += cohesion.Y + separation.Y + alignment.Y
    b.DirVec.Normalize()

    // Update position with scaled speed
    speed := 4.0 // Base speed
    b.Position.X += b.DirVec.X * speed
    b.Position.Y += b.DirVec.Y * speed
}


func (b *Boid) KeepInBounds(w, h float64) {
	if b.Position.X < 0 {
		b.Position.X = w
	} else if b.Position.X > w {
		b.Position.X = 0
	}

	if b.Position.Y < 0 {
		b.Position.Y = h
	} else if b.Position.Y > h {
		b.Position.Y = 0
	}
}
