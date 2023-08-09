package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window 333")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	core.LoadGroundAtlas()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}

func draw() {
	rl.ClearBackground(rl.Black)
	core.DrawTile()
}

func cleanUp() {
	core.TM_CleanUp()
}
