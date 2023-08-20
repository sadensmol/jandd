package editor

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
)

const TileComponentHeight = 120.0

var SelectedTile *core.Tile

type TilesComponent struct {
	selectedAtlas *core.Atlas
	panelScroll   rl.Vector2
}

func (t *TilesComponent) Init() {
	t.panelScroll = rl.Vector2{}
	t.selectedAtlas = &core.TempleAtlas
}

func (t *TilesComponent) Update() {
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		SelectedTile = nil
	}
}

func (t *TilesComponent) Draw(w float32, h float32) {
	panelRec := rl.Rectangle{0, h - TileComponentHeight, w, TileComponentHeight}
	panelContentRec := rl.Rectangle{0, 0, float32(t.selectedAtlas.TotalWidth), float32(t.selectedAtlas.MaxHeight)}

	view := gui.ScrollPanel(panelRec, "Tiles", panelContentRec, &t.panelScroll)

	pad := 30
	gap := 1

	offset := pad

	mPos := rl.GetMousePosition()

	rl.BeginScissorMode(int32(view.X), int32(view.Y), int32(view.Width), int32(view.Height))

	mHov := ""
	for _, v := range t.selectedAtlas.Frames {

		drawPos := rl.Rectangle{X: panelRec.X + t.panelScroll.X + float32(offset), Y: panelRec.Y + t.panelScroll.Y + float32(pad), Width: float32(v.Rect.W), Height: float32(v.Rect.H)}

		rl.DrawTexturePro(t.selectedAtlas.T,
			rl.Rectangle{X: float32(v.Rect.X), Y: float32(v.Rect.Y), Width: float32(v.Rect.W), Height: float32(v.Rect.H)},
			drawPos, rl.Vector2{}, 0, rl.White)

		if rl.CheckCollisionPointRec(mPos, drawPos) {
			rl.DrawRectangleLinesEx(drawPos, 1, rl.Red)
			mHov = v.Name
		}

		offset += v.Rect.W + gap
	}

	rl.EndScissorMode()

	if mHov != "" && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		SelectedTile = &core.Tile{Atlas: t.selectedAtlas, Name: mHov}
		fmt.Printf("selected tile:%v", SelectedTile)
	}

	t.drawAtlasButtons(w, h)

}

func (m *TilesComponent) drawAtlasButtons(w, h float32) {
	pad := 10
	b := gui.Button(rl.Rectangle{X: float32(pad), Y: h - TileComponentHeight, Width: 20, Height: 20}, "T")
	if b {
		m.selectedAtlas = &core.TempleAtlas
	}
	b = gui.Button(rl.Rectangle{X: float32(pad + 20), Y: h - TileComponentHeight, Width: 20, Height: 20}, "G")
	if b {
		m.selectedAtlas = &core.GroundAtlas
	}

}

func (t *TilesComponent) CleanUp() {
}
