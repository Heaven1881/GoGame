package comp

import (
	"math"
	"math/rand"
)

type Position struct {
	X, Y float64
}

func (p *Position) Rand(r int)  {
	p.X = float64(rand.Intn(r))
	p.Y = float64(rand.Intn(r))
}

func (p *Position) DistanceTo(other *Position) float64 {
	var disX = p.X - other.X
	var disY = p.Y - other.Y
	return math.Sqrt(disX*disX + disY*disY)
}

func (p *Position) DirectionTo(other *Position) float64 {
	return math.Atan((other.Y-p.Y)/(other.X-p.X)) / math.Pi * 180
}

type Mover struct {
	Dir   float64 // 角度
	Speed float64 // 速度
}

type AIMoveToTarget struct {
	TargetEntityId uint64
	MaxSpeed float64
}