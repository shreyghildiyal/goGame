package components

type InSystem struct {
	systemId int
	baseComponent
}

func (is *InSystem) SetSystemId(systemId int) {
	is.systemId = systemId
}

func (is *InSystem) GetSystemId() int {
	return is.systemId
}
