package boids

type Position struct {
  X uint64
  Y uint64
}

type UnitVector struct {
  X float32
  Y float32
}

type ForceScalar float32

type Boid struct {
  Position Position
  Direction UnitVector
  Speed ForceScalar
}
