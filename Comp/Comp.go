package Comp

import "math/rand"

type Position struct {
	X, Y float64
}

func (p *Position) Rand(r int)  {
	p.X = float64(rand.Intn(r))
	p.Y = float64(rand.Intn(r))
}

type Mover struct {
	Dir   float64 // 角度
	Speed float64 // 速度
}

type AIMoveToTarget struct {
	TargetEntityId uint64
	MaxSpeed float64
}