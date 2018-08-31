package main

import (
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/key"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/scene"
	"image/color"
)

const Ground collision.Label = 1

func main() {
	//rand.Seed(time.Now().Unix())
	//
	//var sysMgr = sysmgr.New()
	//
	//var mouse = entity.NewSimpleEntity()
	//mouse.CompPosition.Rand(100)
	//
	//var cat = entity.NewSimpleEntity()
	//cat.CompPosition.Rand(100)
	//
	//cat.SetMoveToTarget(mouse, 10)
	//
	//sysMgr.AddEntity(mouse)
	//sysMgr.AddEntity(cat)
	//sysMgr.AddSubSystem(system.SimpleMover{})
	//sysMgr.AddSubSystem(system.AIMoveToTarget{})
	//
	//mouse.Print()
	//for i := 0; i < 10; i++ {
	//	sysMgr.Run()
	//
	//	cat.Print()
	//}

	oak.Add("game",
		func(prevScene string, inData interface{}) {
			char := entities.NewMoving(100, 100, 16, 32,
				render.NewColorBox(16, 32, color.RGBA{R: 255, A: 255}),
				nil, 0, 0)

			render.Draw(char.R)

			char.Speed = physics.NewVector(3, 3)

			char.Bind(func(id int, nothing interface{}) int {
				char := event.GetEntity(id).(*entities.Moving)
				// Move left and right with A and D
				if oak.IsDown(key.A) {
					char.ShiftX(-char.Speed.X())
				}
				if oak.IsDown(key.D) {
					char.ShiftX(char.Speed.X())
				}

				fallSpeed := .1
				oldY := char.Y()
				char.ShiftY(char.Delta.Y())
				hit := collision.HitLabel(char.Space, Ground)
				// If we've moved in y value this frame and in the last frame,
				// we were below what we're trying to hit, we are still falling
				if hit != nil && !(oldY != char.Y() && oldY+char.H > hit.Y()) {
					// Correct our y if we started falling into the ground
					char.SetY(hit.Y() - char.H)
					// Stop falling
					char.Delta.SetY(0)
					// Jump with Space when on the ground
					if oak.IsDown(key.Spacebar) {
						char.Delta.ShiftY(-char.Speed.Y())
					}
				} else {
					// Fall if there's no ground
					char.Delta.ShiftY(fallSpeed)
				}

				return 0
			}, event.Enter)

			ground := entities.NewSolid(0, 400, 500, 20,
				render.NewColorBox(500, 20, color.RGBA{0, 0, 255, 255}),
				nil, 0)
			ground.UpdateLabel(Ground)

			render.Draw(ground.R)
		},
		func() bool {
			return true
		},
		func() (string, *scene.Result) {
			return "game", nil
		},
	)
	oak.Init("game")
}
