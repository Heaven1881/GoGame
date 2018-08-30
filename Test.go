package main

import (
	"./Entity"
	"./System"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var sysMgr = System.NewSysManager()

	var mouse = Entity.NewSimpleEntity()
	mouse.CompPosition.Rand(100)

	var cat = Entity.NewSimpleEntity()
	cat.CompPosition.Rand(100)

	cat.SetMoveTarget(mouse, 10)

	sysMgr.AddEntity(mouse)
	sysMgr.AddEntity(cat)
	sysMgr.AddSystem(System.SimpleMover{})
	sysMgr.AddSystem(System.AIMoveToTarget{})

	mouse.Print()
	for i := 0; i < 10; i++ {
		sysMgr.Run()


		cat.Print()
	}

}

