package system

import "ecs/entity"

type BaseSystem interface {
	tick(sysMgr *SysManager, entity *entity.SimpleEntity)
}

type BaseSystemFunc func(sysMgr *SysManager, entity *entity.SimpleEntity)

func (f BaseSystemFunc) tick(sysMgr *SysManager, entity *entity.SimpleEntity) {
	f(sysMgr, entity)
}

type SysManager struct {
	systems  []BaseSystem
	entities map[uint64]*entity.SimpleEntity
}

func NewSysManager() *SysManager {
	return &SysManager{
		systems:  make([]BaseSystem, 0),
		entities: make(map[uint64]*entity.SimpleEntity, 0),
	}
}

func (sysMgr *SysManager) Run() {
	for _, sys := range sysMgr.systems {
		for _, entity := range sysMgr.entities {
			sys.tick(sysMgr, entity)
		}
	}
}

func (sysMgr *SysManager) AddEntity(entity *entity.SimpleEntity) {
	sysMgr.entities[entity.EntityId()] = entity
}

func (sysMgr *SysManager) AddSystem(system BaseSystem) {
	sysMgr.systems = append(sysMgr.systems, system)
}
