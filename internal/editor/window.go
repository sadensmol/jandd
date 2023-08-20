package editor

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
)

const (
	winWidth  = 1024
	WinHeight = 768
)

var components []core.IComponent

func InitWindow() {
	rl.InitWindow(winWidth, WinHeight, "editor")
	rl.SetWindowPosition(100, 50)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	components = append(components, &TilesComponent{}, &MapComponent{}, &SettingsComponent{})
	for _, c := range components {
		c.Init()
	}
	core.LoadAtlases()

	for !rl.WindowShouldClose() {
		UpdateAndDrawWindow()
	}
}

func UpdateAndDrawWindow() {
	for _, c := range components {
		c.Update()
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for _, c := range components {
		c.Draw(winWidth, WinHeight)
	}
	rl.EndDrawing()
}
