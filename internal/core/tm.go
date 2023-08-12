package core

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Atlas *Atlas
	Name  string
}

type Atlas struct {
	T          rl.Texture2D
	Frames     `json:"frames"`
	TotalWidth int
	MaxHeight  int
}

type Frame struct {
	Name string `json:"name"`
	Rect Rect   `json:"rect"`
}

type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Frames struct {
	Frames []Frame `json:"frames"`
}

var GroundAtlas Atlas
var TempleAtlas Atlas

func (a Atlas) FrameRect(name string) rl.Rectangle {
	for _, f := range a.Frames.Frames {
		if f.Name == name {
			return rl.Rectangle{X: float32(f.Rect.X), Y: float32(f.Rect.Y), Width: float32(f.Rect.W), Height: float32(f.Rect.H)}
		}
	}
	println("not found!")
	return rl.Rectangle{}
}

func TM_CleanUp() {
	rl.UnloadTexture(GroundAtlas.T)
}

func LoadAtlases() {
	GroundAtlas = loadAtlas("ground")
	TempleAtlas = loadAtlas("temple")
}

func loadAtlas(name string) Atlas {
	t := rl.LoadTexture(fmt.Sprintf("./assets/%s.png", name))

	data, err := os.ReadFile(fmt.Sprintf("./assets/%s.json", name))

	for _, v := range data {
		fmt.Printf("%c", v)
	}

	if err != nil {
		panic(err)
	}

	var frames Frames
	err = json.Unmarshal(data, &frames)
	if err != nil {
		panic(err)
	}

	tw := 0
	mh := 0
	for _, v := range frames.Frames {
		tw += v.Rect.W
		mh = int(math.Max(float64(mh), float64(v.Rect.H)))
	}

	return Atlas{T: t, Frames: frames, TotalWidth: tw, MaxHeight: mh}
}

func DrawTile() {

}
