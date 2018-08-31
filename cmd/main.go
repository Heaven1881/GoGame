package main

import (
	"ecs/entity"
	"ecs/system"
	"math/rand"
	"time"
)

import _ "github.com/oakmound/oak"

func main() {
	rand.Seed(time.Now().Unix())

	var sysMgr = system.NewSysManager()

	var mouse = entity.NewSimpleEntity()
	mouse.CompPosition.Rand(100)

	var cat = entity.NewSimpleEntity()
	cat.CompPosition.Rand(100)

	cat.SetMoveToTarget(mouse, 10)

	sysMgr.AddEntity(mouse)
	sysMgr.AddEntity(cat)
	sysMgr.AddSystem(system.SimpleMover{})
	sysMgr.AddSystem(system.AIMoveToTarget{})

	mouse.Print()
	for i := 0; i < 10; i++ {
		sysMgr.Run()


		cat.Print()
	}

}

