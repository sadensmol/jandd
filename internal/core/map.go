package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Map struct {
	Elements []MapElement
}

type MapElement struct {
	Tile  *Tile
	Pos   rl.Vector2
	Layer int
}

func (m *Map) AddElement(e MapElement) {
	m.Elements = append(m.Elements, e)
}
