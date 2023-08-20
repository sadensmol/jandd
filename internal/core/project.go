package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Project struct {
	Version int
	Map     ProjectMap
}

type ProjectMapElement struct {
	Name  string
	Pos   rl.Vector2
	Layer int
}

type ProjectMap struct {
	Elements []ProjectMapElement
}

func ProjectMapToMap(pm *ProjectMap) Map {
	m := Map{}
	for _, e := range pm.Elements {
		t := FindTile(e.Name)
		if t != nil {
			m.AddElement(MapElement{Tile: t, Pos: e.Pos, Layer: e.Layer})
		}

	}
	return m
}
func MapToProjectMap(m *Map) ProjectMap {
	pm := ProjectMap{}
	for _, e := range m.Elements {
		pm.Elements = append(pm.Elements, ProjectMapElement{
			Name:  e.Tile.Name,
			Pos:   e.Pos,
			Layer: e.Layer,
		})
	}
	return pm
}
