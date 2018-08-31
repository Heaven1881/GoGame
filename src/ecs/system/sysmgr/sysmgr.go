package sysmgr

import "ecs/entity"

type BaseSystem interface {
	Tick(sysMgr *SysManager, entity *entity.SimpleEntity)
}

type BaseSystemFunc func(sysMgr *SysManager, entity *entity.SimpleEntity)

func (f BaseSystemFunc) tick(sysMgr *SysManager, entity *entity.SimpleEntity) {
	f(sysMgr, entity)
}

type SysManager struct {
	systems  []BaseSystem
	entities map[uint64]*entity.SimpleEntity
}

func New() *SysManager {
	return &SysManager{
		systems:  make([]BaseSystem, 0),
		entities: make(map[uint64]*entity.SimpleEntity, 0),
	}
}

func (sysMgr *SysManager) Run() {
	for _, sys := range sysMgr.systems {
		for _, entt := range sysMgr.entities {
			sys.Tick(sysMgr, entt)
		}
	}
}

func (sysMgr *SysManager) AddEntity(entity *entity.SimpleEntity) {
	sysMgr.entities[entity.EntityId()] = entity
}

func (sysMgr *SysManager) GetEntity(id uint64) *entity.SimpleEntity  {
	return sysMgr.entities[id]
}

func (sysMgr *SysManager) AddSubSystem(system BaseSystem) {
	sysMgr.systems = append(sysMgr.systems, system)
}
