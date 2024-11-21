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

func inViewRange(origin, other *Boid, viewAngle, viewDistance float64) bool {
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
        if inViewRange(b, &neighbor, viewAngle, viewDistance) {
            neighborCount++

            // Cohesion: Average position
            cohesion.X += neighbor.Position.X
            cohesion.Y += neighbor.Position.Y

            // Separation: Avoidance
            delta := Vector{
                X: b.Position.X - neighbor.Position.X,
                Y: b.Position.Y - neighbor.Position.Y,
            }
            distance := math.Sqrt(delta.X*delta.X + delta.Y*delta.Y)
            if distance != 0 {
                separation.X += delta.X / distance
                separation.Y += delta.Y / distance
            }

            // Alignment: Average direction
            alignment.X += neighbor.DirVec.X
            alignment.Y += neighbor.DirVec.Y
        }
    }

    if neighborCount > 0 {
        // Cohesion
        cohesion.X /= float64(neighborCount)
        cohesion.Y /= float64(neighborCount)
        cohesion.X -= b.Position.X
        cohesion.Y -= b.Position.Y
        cohesion.Normalize()

        // Separation
        separation.Normalize()

        // Alignment
        alignment.X /= float64(neighborCount)
        alignment.Y /= float64(neighborCount)
        alignment.Normalize()
    }

    // Combine forces
    b.DirVec.add(&cohesion)
    b.DirVec.add(&separation)
    b.DirVec.add(&alignment)
    b.DirVec.Normalize()

    // Update position
    b.Position.X += b.DirVec.X
    b.Position.Y += b.DirVec.Y
}

