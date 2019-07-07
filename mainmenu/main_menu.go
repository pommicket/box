package mainmenu

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/eng/widgets"
	"github.com/pommicket/box/state"
)

var start widgets.Button
var levelEditor widgets.Button
var exit widgets.Button
var nextState state.State
var shown bool

func Load() {
	start.LoadAll("start.bmp")
	start.Scale = 4
	start.Pos.Align = widgets.TOP_MIDDLE
	start.OnClick = func() {
		nextState = state.LEVEL_SELECT
	}
	levelEditor.LoadAll("level_editor.bmp")
	levelEditor.Scale = 4
	levelEditor.Pos.SetParent(&start.Pos, widgets.TOP_MIDDLE, widgets.BOTTOM_MIDDLE)
	levelEditor.Pos.Move(0, 8)
	levelEditor.OnClick = func() {
		nextState = state.LEVEL_EDITOR
	}
	exit.LoadAll("exit.bmp")
	exit.Scale = 4
	exit.Pos.SetParent(&levelEditor.Pos, widgets.TOP_MIDDLE, widgets.BOTTOM_MIDDLE)
	exit.Pos.Move(0, 8)
	exit.OnClick = func() {
		nextState = state.EXIT
	}
	eng.OnKeyUp(keyUp)
}

func Show() {
	start.Show()
	levelEditor.Show()
	exit.Show()
	nextState = state.MAIN_MENU
	shown = true
}

func Hide() {
	start.Hide()
	levelEditor.Hide()
	exit.Hide()
	shown = false
}

func Render() state.State {
	start.Pos.Move(eng.Width()/2, 8)
	eng.SetColor(common.Color1)
	eng.Clear()
	return nextState
}

func keyUp(key int) {
	if !shown {
		return
	}
	switch key {
	case eng.KEY_ESCAPE:
		nextState = state.EXIT
	}
}
