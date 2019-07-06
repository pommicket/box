package mainmenu

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/eng/widgets"
	"github.com/pommicket/box/state"
)

var start widgets.Button
var levelEditor widgets.Button
var nextState state.State

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
}

func Show() {
	start.Show()
	levelEditor.Show()
	nextState = state.MAIN_MENU
}

func Hide() {
	start.Hide()
	levelEditor.Hide()
}

func Render() state.State {
	start.Pos.Move(eng.Width()/2, 8)
	eng.SetColor(common.Color1)
	eng.Clear()
	return nextState
}
