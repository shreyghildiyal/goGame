package components

type Name struct {
	Name     string
	entityId int
}

func (n Name) GetEnitityId() int {
	return n.entityId
}
