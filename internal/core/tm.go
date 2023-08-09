package core

import (
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Atlas struct {
	T      rl.Texture2D
	Frames `json:"frames"`
}

type Frame struct {
	Filename         string `json:"filename"`
	Frame            Rect   `json:"frame"`
	Rotated          bool   `json:"rotated"`
	Trimmed          bool   `json:"trimmed"`
	SpriteSourceSize Rect   `json:"spriteSourceSize"`
	SourceSize       Size   `json:"sourceSize"`
}

type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Frames struct {
	Frames []Frame `json:"frames"`
}

var A Atlas

func TM_CleanUp() {
	rl.UnloadTexture(A.T)
}

func LoadGroundAtlas() {
	t := rl.LoadTexture("./assets/ground.png")

	data, err := os.ReadFile("./assets/ground.json")
	if err != nil {
		panic(err)
	}

	var frames Frames
	err = json.Unmarshal(data, &frames)
	if err != nil {
		panic(err)
	}

	A = Atlas{T: t, Frames: frames}
}

func DrawTile() {
	rl.DrawTexturePro(A.T, rl.Rectangle{X: 0, Y: 0, Width: float32(A.Frames.Frames[0].SourceSize.W), Height: float32(A.Frames.Frames[0].SourceSize.H)}, rl.Rectangle{X: 0, Y: 0, Width: 32, Height: 32}, rl.Vector2{}, 0, rl.White)
}
