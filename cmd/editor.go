package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
	"github.com/sadensmol/jandd/internal/editor"
)

const (
	winWidth  = 1024
	winHeight = 768
)

var components []core.IComponent

func Init() {
	components = append(components, &editor.TilesComponent{}, &editor.MapComponent{})

	for _, c := range components {
		c.Init()
	}

	core.LoadAtlases()

}

func UpdateAndDraw() {

	for _, c := range components {
		c.Update()
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, c := range components {
		c.Draw(winWidth, winHeight)
	}

	rl.EndDrawing()
}

func CleanUp() {
	core.TM_CleanUp()
}

func main() {
	rl.InitWindow(winWidth, winHeight, "editor")
	rl.SetWindowPosition(100, 50)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	Init()

	for !rl.WindowShouldClose() {
		UpdateAndDraw()
	}
	CleanUp()
}
