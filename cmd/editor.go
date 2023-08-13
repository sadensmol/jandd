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

	// test() // test shaders code is located below

	Init()

	for !rl.WindowShouldClose() {
		UpdateAndDraw()
	}
	CleanUp()
}

var tx rl.Texture2D
var rt rl.Texture2D

// move to heap
func lt() {
	r := rl.LoadRenderTexture(600, 600)

	rl.BeginTextureMode(r)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))
	rl.DrawCircle(250, 300, 160, rl.NewColor(0, 0, 0, 255))
	rl.EndTextureMode()
	rt = r.Texture

	tx = rl.LoadTexture("test.png")
}

func test() {
	lt()
	sh := rl.LoadShader("", "./shaders/fmask.fs")

	sl := rl.GetShaderLocation(sh, "texture1")

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginShaderMode(sh)
		rl.SetShaderValueTexture(sh, sl, rt)   // texture1
		rl.DrawTexture(tx, 0.0, 0.0, rl.White) // texture0
		rl.EndShaderMode()

		rl.EndDrawing()
	}

	rl.UnloadShader(sh)
	rl.UnloadTexture(tx)
	rl.UnloadTexture(rt)
	rl.CloseWindow()
}
