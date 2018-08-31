package main

import (
	"ecs/entity"
	"ecs/system"
	"ecs/system/sysmgr"
	"math/rand"
	"time"
)

import _ "github.com/oakmound/oak"

func main() {
	rand.Seed(time.Now().Unix())

	var sysMgr = sysmgr.New()

	var mouse = entity.NewSimpleEntity()
	mouse.CompPosition.Rand(100)

	var cat = entity.NewSimpleEntity()
	cat.CompPosition.Rand(100)

	cat.SetMoveToTarget(mouse, 10)

	sysMgr.AddEntity(mouse)
	sysMgr.AddEntity(cat)
	sysMgr.AddSubSystem(system.SimpleMover{})
	sysMgr.AddSubSystem(system.AIMoveToTarget{})

	mouse.Print()
	for i := 0; i < 10; i++ {
		sysMgr.Run()


		cat.Print()
	}

}

