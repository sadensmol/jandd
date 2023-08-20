package editor

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sadensmol/jandd/internal/core"
	"os"
)

const projectsDir = "./projects/"

type SettingsComponent struct {
	isOpen bool
}

func (m *SettingsComponent) Init() {
}

func (m *SettingsComponent) Update() {
}

var projectFileName = "test"

func (m *SettingsComponent) Draw(w float32, h float32) {

	gui.SetState(gui.STATE_NORMAL)
	if m.isOpen {
		gui.SetState(gui.STATE_PRESSED)
	}

	settingsButton := gui.Button(rl.Rectangle{X: w - 40, Y: 20, Width: 20, Height: 20}, "S")
	if settingsButton {
		m.isOpen = !m.isOpen
	}

	if m.isOpen {
		var pad float32 = 20
		var width float32 = 200
		left := w - width - pad

		gui.SetState(gui.STATE_NORMAL)

		gui.TextBox(rl.Rectangle{X: left, Y: 60, Width: width, Height: 20}, &projectFileName, 20, true)
		cstr := C.CString(projectFileName)
		projectFileName = C.GoString(cstr)

		loadButton := gui.Button(rl.Rectangle{X: left, Y: 60 + pad, Width: width, Height: 20}, "Load")
		_ = loadButton
		if loadButton {
			loadProject()
			m.isOpen = false
		}
		saveButton := gui.Button(rl.Rectangle{X: left, Y: 60 + pad*2, Width: width, Height: 20}, "Save")
		if saveButton {
			saveProject()
			m.isOpen = false
		}

		resetButton := gui.Button(rl.Rectangle{X: left, Y: 60 + pad*3, Width: width, Height: 20}, "Reset")
		if resetButton {
			// resetMap()
		}
	}
}

func saveProject() {
	f, err := os.Create(fmt.Sprintf("%s%s.jadp", projectsDir, projectFileName))
	if err != nil {
		rl.TraceLog(rl.LogError, "cannot save project %w", err)
		return
	}
	defer f.Close()

	p := core.Project{Version: 1, Map: core.MapToProjectMap(&Map)}

	jd, err := json.Marshal(p)
	if err != nil {
		rl.TraceLog(rl.LogError, "%w", err)
		return
	}
	_, err = f.Write(jd)
	if err != nil {
		rl.TraceLog(rl.LogError, "%w", err)
		return
	}
}
func loadProject() {
	jd, err := os.ReadFile(fmt.Sprintf("%s%s.jadp", projectsDir, projectFileName))
	if err != nil {
		rl.TraceLog(rl.LogError, "Cannot read file :", err)
		return
	}

	p := core.Project{}
	err = json.Unmarshal(jd, &p)
	if err != nil {
		rl.TraceLog(rl.LogError, "%w", err)
		return
	}

	Map = core.ProjectMapToMap(&p.Map)
}

func (m *SettingsComponent) CleanUp() {

}
