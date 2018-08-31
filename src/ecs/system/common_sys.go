package system

import (
	"ecs/entity"
	"ecs/system/sysmgr"
	"math"
)

type SimpleMover struct {
}

func (SimpleMover) Tick(sys *sysmgr.SysManager, e *entity.SimpleEntity) {
	if e.CompMover != nil &&
		e.CompPosition != nil {
		e.CompPosition.X += math.Cos(math.Pi*e.CompMover.Dir/180.0) * e.CompMover.Speed
		e.CompPosition.Y += math.Sin(math.Pi*e.CompMover.Dir/180.0) * e.CompMover.Speed
	}
}

func DistanceBetween(e1 *entity.SimpleEntity, e2 *entity.SimpleEntity) float64 {
	if e1.CompPosition != nil && e2.CompPosition != nil {
		return e1.CompPosition.DistTo(e2.CompPosition)
	} else {
		// 没有位置的Entity之间，距离为无穷大
		return math.Inf(1)
	}
}

type AIMoveToTarget struct {
}

func (AIMoveToTarget) Tick(sys *sysmgr.SysManager, e *entity.SimpleEntity) {
	if e.CompPosition != nil &&
		e.CompMover != nil &&
		e.CompAIMoveToTarget != nil {
		var target = sys.GetEntity(e.CompAIMoveToTarget.TargetEntityId)

		if target != nil {
			var dist = e.CompPosition.DistTo(target.CompPosition)

			if dist > 0 {
				e.CompMover.Dir = e.CompPosition.DirTo(target.CompPosition)
				e.CompMover.Speed = math.Min(dist, e.CompAIMoveToTarget.MaxSpeed)
			} else {
				e.CompMover.Dir = 0
				e.CompMover.Speed = 0
			}
		}
	}
}
