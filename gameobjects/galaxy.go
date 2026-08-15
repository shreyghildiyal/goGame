package gameobjects

type Galaxy struct {
	Stars      map[int]Star
	Planets    map[int]Planet
	StarSystem map[int]StarSystem
}
