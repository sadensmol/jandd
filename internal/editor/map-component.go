package editor

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
)

const (
	MapWidth  = 1000
	MapHeight = 640
)

const (
	MapLayerBackground = iota
	MapLayerForeground
	MapLayerForegroundMask
	MapLayerIcons
)

type MapComponent struct {
	selectedLayer int
	elements      []MapElement
}

type MapElement struct {
	Tile  *core.Tile
	Pos   rl.Vector2
	Layer int
}

var maskShader rl.Shader
var fgTexture rl.RenderTexture2D
var fgMaskTexture rl.RenderTexture2D
var fgMaskShLoc int32

func (m *MapComponent) Init() {
	maskShader = rl.LoadShader("", "./shaders/fmask.fs")
	fgTexture = rl.LoadRenderTexture(MapWidth, MapHeight)
	fgMaskTexture = rl.LoadRenderTexture(MapWidth, MapHeight)
	fgMaskShLoc = rl.GetShaderLocation(maskShader, "texture1")
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

	gui.SetState(gui.STATE_NORMAL)
	if m.selectedLayer == MapLayerBackground {
		gui.SetState(gui.STATE_PRESSED)
	}

	bg := gui.Button(rl.Rectangle{X: float32(pad + 20), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "B")
	if bg {
		m.selectedLayer = MapLayerBackground
	}

	gui.SetState(gui.STATE_NORMAL)
	if m.selectedLayer == MapLayerForeground {
		gui.SetState(gui.STATE_PRESSED)
	}

	fg := gui.Button(rl.Rectangle{X: float32(pad + 40), Y: h - TileComponentHeight - 50, Width: 20, Height: 20}, "F")
	if fg {
		m.selectedLayer = MapLayerForeground
	}

	gui.SetState(gui.STATE_NORMAL)
	if m.selectedLayer == MapLayerForegroundMask {
		gui.SetState(gui.STATE_PRESSED)
	}

	fgm := gui.Button(rl.Rectangle{X: float32(pad + 40), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "M")
	if fgm {
		m.selectedLayer = MapLayerForegroundMask
	}

	gui.SetState(gui.STATE_NORMAL)
	if m.selectedLayer == MapLayerIcons {
		gui.SetState(gui.STATE_PRESSED)
	}

	ic := gui.Button(rl.Rectangle{X: float32(pad + 60), Y: h - TileComponentHeight - 30, Width: 20, Height: 20}, "I")
	if ic {
		m.selectedLayer = MapLayerIcons
	}

	gui.SetState(gui.STATE_NORMAL)
}

// todo draw into a texture and put texture into scroll pane?
func (m *MapComponent) drawMapElements() {
	layers := []int{MapLayerBackground, MapLayerForeground, MapLayerForegroundMask, MapLayerIcons}

	for _, l := range layers {
		if l > m.selectedLayer {
			break
		}

		var lTex rl.RenderTexture2D
		txRecFlipped := rl.Rectangle{X: 0, Y: 0, Width: MapWidth, Height: -MapHeight}

		// we draw fg into texture
		if l == MapLayerForeground {
			lTex = fgTexture
		}

		// we draw fg mask into texture first
		if l == MapLayerForegroundMask {
			lTex = fgMaskTexture
		}

		if l == MapLayerForeground || l == MapLayerForegroundMask {
			rl.BeginTextureMode(lTex)
			rl.ClearBackground(rl.Blank)
		}

		for _, v := range m.elements {
			if v.Layer != l { // draw elements only from current layer
				continue
			}
			f := v.Tile.Atlas.FrameRect(v.Tile.Name)
			rl.DrawTexturePro(v.Tile.Atlas.T, f, rl.Rectangle{v.Pos.X, v.Pos.Y, f.Width, f.Height}, rl.Vector2{}, 0, rl.White)
		}

		// we draw everything on fg and fgmask layer if one of them are chosen
		// if next layer is chosen then we use masking
		if l == MapLayerForeground || l == MapLayerForegroundMask {
			rl.EndTextureMode()
			if m.selectedLayer <= MapLayerForegroundMask {
				rl.DrawTextureRec(lTex.Texture, txRecFlipped, rl.Vector2{0, 0}, rl.White)
			} else {
				if l == MapLayerForegroundMask {
					rl.BeginShaderMode(maskShader)
					rl.SetShaderValueTexture(maskShader, fgMaskShLoc, fgMaskTexture.Texture)       // texture1
					rl.DrawTextureRec(fgTexture.Texture, txRecFlipped, rl.Vector2{0, 0}, rl.White) // texture0
					rl.EndShaderMode()
				}
			}
		}
	}
}

func (m *MapComponent) CleanUp() {
	rl.UnloadShader(maskShader)
	rl.UnloadRenderTexture(fgTexture)
	rl.UnloadRenderTexture(fgMaskTexture)
}
