package drawing

import "fmt"

type DrawableRegistry struct {
	galaxyDrawables map[string][]Drawable
	systemDrawables map[string][]Drawable
}

func NewDrawableRegistry() DrawableRegistry {
	dr := DrawableRegistry{
		galaxyDrawables: map[string][]Drawable{},
		systemDrawables: map[string][]Drawable{},
	}

	return dr
}

func (dr *DrawableRegistry) AddGalaxyDrawables(id string, drawables []Drawable) {
	dr.galaxyDrawables[id] = drawables
}

func (dr *DrawableRegistry) AddSystemDrawables(id string, drawables []Drawable) {
	dr.systemDrawables[id] = drawables
}

func (dr *DrawableRegistry) RemoveGalaxyDrawables(id string) {
	delete(dr.galaxyDrawables, id)
}

func (dr *DrawableRegistry) RemoveSystemDrawables(id string) {
	delete(dr.systemDrawables, id)
}

func (dr *DrawableRegistry) GetGalaxyDrawables(id string) ([]Drawable, error) {
	if drawables, found := dr.galaxyDrawables[id]; found {
		return drawables, nil
	} else {
		return nil, fmt.Errorf("No Galaxy drawables registered for %s", id)
	}
}

func (dr *DrawableRegistry) GetSystemDrawables(id string) ([]Drawable, error) {
	if drawables, found := dr.systemDrawables[id]; found {
		return drawables, nil
	} else {
		return nil, fmt.Errorf("No System drawables registered for %s", id)
	}
}
