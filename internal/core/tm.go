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
	Frames     []Frame `json:"frames"`
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

var GroundAtlas Atlas
var TempleAtlas Atlas

func FindTile(name string) *Tile {
	if f := GroundAtlas.FindFrame(name); f != nil {
		return &Tile{Atlas: &GroundAtlas, Name: name}
	}

	if f := TempleAtlas.FindFrame(name); f != nil {
		return &Tile{Atlas: &TempleAtlas, Name: name}
	}

	return nil
}

func (a Atlas) FindFrame(name string) *Frame {
	for _, f := range a.Frames {
		if f.Name == name {
			return &f
		}
	}
	return nil
}
func (a Atlas) FrameRect(name string) rl.Rectangle {
	for _, f := range a.Frames {
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

	a := Atlas{T: t, Frames: make([]Frame, 0)}

	data, err := os.ReadFile(fmt.Sprintf("./assets/%s.json", name))

	if err != nil {
		panic(err)
	}

	type Frames struct {
		Frames []Frame `json:"frames"`
	}
	f := Frames{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		panic(err)
	}

	a.Frames = f.Frames

	tw := 0
	mh := 0
	for _, v := range a.Frames {
		tw += v.Rect.W
		mh = int(math.Max(float64(mh), float64(v.Rect.H)))
	}

	a.TotalWidth = tw
	a.MaxHeight = mh

	return a
}

func DrawTile() {

}
