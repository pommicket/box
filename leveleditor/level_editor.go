package leveleditor

import (
	"fmt"
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/levels"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
)

var kindsForKeys = map[int]objects.ObjectKind{
	eng.KEY_b: objects.BOX,
	eng.KEY_e: objects.ENEMY,
	eng.KEY_s: objects.SPIKE,
	eng.KEY_l: objects.CONVEYOR_LEFT,
	eng.KEY_r: objects.CONVEYOR_RIGHT,
	eng.KEY_p: objects.PORTAL,
	eng.KEY_g: objects.GOAL,
}

var currentlyPlacing objects.ObjectKind
var saveFile string
var nextState state.State
var shown bool

func Load() {
	eng.OnKeyUp(keyUp)
	eng.OnMouseMove(mouseMove)
}

func Render() state.State {
	eng.SetColor(common.Color1)
	eng.Clear()
	objects.RenderAll(true)
	mouseX, mouseY := eng.MousePos()
	mouseMove(mouseX, mouseY)
	if currentlyPlacing != objects.NONE {
		tileX, tileY := objects.PixelToTile(mouseX, mouseY)
		var obj objects.Object
		obj.Set(currentlyPlacing, false)
		obj.X = tileX
		obj.Y = tileY
		obj.RenderGhost()
	}

	return nextState
}

func keyUp(keyCode int) {
	if !shown {
		return
	}
	switch keyCode {
	case eng.KEY_ESCAPE:
		if currentlyPlacing == objects.NONE {
			nextState = state.MAIN_MENU
		}
		currentlyPlacing = objects.NONE
		return
	case eng.KEY_s:
		if eng.IsCtrl() {
			if eng.IsShift() || saveFile == "" {
				fmt.Print("Save as? ")
				fmt.Scanln(&saveFile)
			}
			if saveFile != "" {
				levels.Save(saveFile)
			}
			return
		}
	case eng.KEY_o:
		if eng.IsCtrl() {
			fmt.Print("Open? ")
			fmt.Scanln(&saveFile)
			if saveFile != "" {
				levels.Load(saveFile)
			}
		}
	}
	objkind, ok := kindsForKeys[keyCode]
	if !ok {
		return
	}
	currentlyPlacing = objkind
}

func mouseMove(x, y int) {
	if !shown {
		return
	}
	if eng.IsMouseDown(eng.MOUSE_LEFT) {
		tileX, tileY := objects.PixelToTile(x, y)
		objects.At(tileX, tileY).Set(currentlyPlacing, true)
	}
}

func Show() {
	nextState = state.LEVEL_EDITOR
	objects.ClearAll()
	shown = true
}

func Hide() {
	shown = false
}
