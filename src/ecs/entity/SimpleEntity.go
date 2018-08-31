package entity

import (
	"ecs/comp"
	"fmt"
)

type SimpleEntity struct {
	entityId           uint64
	CompPosition       *comp.Position
	CompMover          *comp.Mover
	CompAIMoveToTarget *comp.AIMoveToTarget
}

func (entity *SimpleEntity) EntityId() uint64 {
	return entity.entityId
}

var nextEntityId uint64 = 0

func NewSimpleEntity() *SimpleEntity {
	nextEntityId += 1
	return &SimpleEntity{
		entityId:     nextEntityId,
		CompMover:    new(comp.Mover),
		CompPosition: new(comp.Position),
	}
}

func (entity *SimpleEntity) Print() {
	fmt.Printf("=== ID: %d\n", entity.entityId)
	fmt.Printf("  > %+v\n", *entity.CompPosition)
	fmt.Printf("  > %+v\n", *entity.CompMover)
}

func (entity *SimpleEntity) SetMoveToTarget(target *SimpleEntity, maxSpeed float64) {
	if target != nil {
		entity.CompAIMoveToTarget = &comp.AIMoveToTarget{
			TargetEntityId: target.EntityId(),
			MaxSpeed:       maxSpeed,
		}
	} else {
		entity.CompAIMoveToTarget = nil
	}

}
