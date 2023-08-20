package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
	"github.com/sadensmol/jandd/internal/editor"
)

func Init() {
}

func CleanUp() {
	core.TM_CleanUp()
}

func main() {

	editor.InitWindow()

	// test() // test shaders code is located below
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
