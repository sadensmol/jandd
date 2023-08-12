package editor

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
)

const (
	MapWidth  = 1000
	MapHeight = 1000
)

const (
	MapLayerBackground = iota
	MapLayerForeground
	MapLayerForegroundMask
	MapLayerIcons
)

type MapComponent struct {
	showAllLayers bool
	selectedLayer int
	elements      []MapElement
}

type MapElement struct {
	Tile  *core.Tile
	Pos   rl.Vector2
	Layer int
}

func (m *MapComponent) Init() {
}

func (m *MapComponent) Update() {
}

func (m *MapComponent) Draw(w float32, h float32) {
	m.drawMapElements()

	mPos := rl.GetMousePosition()

	mapRect := rl.Rectangle{X: 0, Y: 0, Width: w, Height: h - TileComponentHeight}

	if rl.CheckCollisionPointRec(mPos, mapRect) {

		if SelectedTile != nil {
			f := SelectedTile.Atlas.FrameRect(SelectedTile.Name)
			rl.DrawTexturePro(SelectedTile.Atlas.T, f, rl.Rectangle{mPos.X, mPos.Y, f.Width, f.Height}, rl.Vector2{}, 0, rl.White)

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				m.elements = append(m.elements, MapElement{Tile: SelectedTile, Pos: mPos, Layer: m.selectedLayer})
			}
		}

	}

	m.drawLayerButtons(w, h)
}

func (m *MapComponent) drawLayerButtons(w, h float32) {
	pad := 10
	all := gui.Button(rl.Rectangle{X: float32(pad), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "A")
	if all {
		m.showAllLayers = !m.showAllLayers
	}
	bg := gui.Button(rl.Rectangle{X: float32(pad + 20), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "B")
	if bg {
		m.selectedLayer = MapLayerBackground
	}
	fg := gui.Button(rl.Rectangle{X: float32(pad + 40), Y: h - TileComponentHeight - 50, Width: 20, Height: 20}, "F")
	if fg {
		m.selectedLayer = MapLayerForeground
	}
	fgm := gui.Button(rl.Rectangle{X: float32(pad + 40), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "M")
	if fgm {
		m.selectedLayer = MapLayerForegroundMask
	}

	ic := gui.Button(rl.Rectangle{X: float32(pad + 60), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "I")
	if ic {
		m.selectedLayer = MapLayerIcons
	}

}

// todo draw into a texture and put texture into scroll pane?
func (m *MapComponent) drawMapElements() {

	layers := []int{MapLayerBackground, MapLayerForeground, MapLayerForegroundMask, MapLayerIcons}

	for _, l := range layers {

		if !m.showAllLayers && l > m.selectedLayer {
			continue
		}

		for _, v := range m.elements {
			if v.Layer != l {
				continue
			}
			f := v.Tile.Atlas.FrameRect(v.Tile.Name)
			rl.DrawTexturePro(v.Tile.Atlas.T, f, rl.Rectangle{v.Pos.X, v.Pos.Y, f.Width, f.Height}, rl.Vector2{}, 0, rl.White)
		}
	}
}

func (m *MapComponent) CleanUp() {
}
