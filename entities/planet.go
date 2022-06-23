package entities

import "github.com/shreyghildiyal/goGame/constants"

type Planet struct {
	Id              int
	ComponentsMap   map[constants.ComponentTypeName][]int
	EntityType      constants.EntityTypeName
	markedForDelete bool
}

// func NewPlanet() Planet {
// 	planet := Planet{}

// 	return planet
// }

func (p *Planet) GetId() int {
	return p.Id
}

func (p *Planet) SetId(newId int) {
	p.Id = newId
}

func (p *Planet) GetComponentIds(componentType constants.ComponentTypeName) ([]int, bool) {
	ids, found := p.ComponentsMap[componentType]
	return ids, found
}

func (p *Planet) GetComponentsMap() map[constants.ComponentTypeName][]int {
	return p.ComponentsMap
}

func (p *Planet) AddComponent(componentType constants.ComponentTypeName, componentId int) {
	if p.ComponentsMap == nil {
		p.ComponentsMap = map[constants.ComponentTypeName][]int{}
	}
	if _, found := p.ComponentsMap[componentType]; !found {
		p.ComponentsMap[componentType] = []int{}
	}
	p.ComponentsMap[componentType] = append(p.ComponentsMap[componentType], componentId)
}

func (p *Planet) GetEntityType() constants.EntityTypeName {
	return p.EntityType
}

func (p *Planet) MarkForDelete() {
	p.markedForDelete = true
}

func (p *Planet) IsMarkedForDelete() bool {
	return p.markedForDelete
}
