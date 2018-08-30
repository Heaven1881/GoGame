package System

import (
	"../Entity"
	"math"
)

type SimpleMover struct {
}

func (SimpleMover) tick(sys *SysManager, e *Entity.SimpleEntity) {
	if e.CompMover != nil &&
		e.CompPosition != nil {
		e.CompPosition.X += math.Cos(math.Pi*e.CompMover.Dir/180.0) * e.CompMover.Speed
		e.CompPosition.Y += math.Sin(math.Pi*e.CompMover.Dir/180.0) * e.CompMover.Speed
	}
}

func distance(e1 *Entity.SimpleEntity, e2 *Entity.SimpleEntity) float64 {
	if e1.CompPosition != nil && e2.CompPosition != nil {
		var disX = e1.CompPosition.X - e2.CompPosition.X
		var disY = e1.CompPosition.Y - e2.CompPosition.Y
		return math.Sqrt(disX*disX + disY*disY)
	} else {
		// 没有位置的Entity之间，距离为无穷大
		return math.Inf(1)
	}
}

type AIMoveToTarget struct {
}

func (AIMoveToTarget) tick(sys *SysManager, e *Entity.SimpleEntity) {
	if e.CompPosition != nil &&
		e.CompMover != nil &&
		e.CompAIMoveToTarget != nil {
		var target = sys.entities[e.CompAIMoveToTarget.TargetEntityId]

		if target != nil {
			var dist = distance(target, e)

			if dist > 0 {

				e.CompMover.Dir = math.Atan((target.CompPosition.Y-e.CompPosition.Y)/(target.CompPosition.X-e.CompPosition.X)) / math.Pi * 180
				e.CompMover.Speed = math.Min(dist, e.CompAIMoveToTarget.MaxSpeed)
			} else {
				e.CompMover.Dir = 0
				e.CompMover.Speed = 0
			}
		}
	}
}