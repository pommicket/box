package mainmenu

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/eng/widgets"
	"github.com/pommicket/box/state"
)

var start, levelEditor, exit, inv1, inv2 widgets.Button
var nextState state.State
var shown, c, h bool
var sprC, sprH eng.Sprite

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
	inv1.LoadAll("invisible.bmp")
	inv1.Scale = 4
	inv1.Pos.SetParent(&exit.Pos, widgets.TOP_MIDDLE, widgets.BOTTOM_MIDDLE)
	inv1.Pos.Move(0, 8)
	inv1.OnClick = func() {
		h = !h
	}
	inv2.LoadAll("invisible.bmp")
	inv2.Scale = 4
	inv2.Pos.SetParent(&inv1.Pos, widgets.TOP_MIDDLE, widgets.BOTTOM_MIDDLE)
	inv2.Pos.Move(0, 8)
	inv2.OnClick = func() {
		c = !c
	}
	sprC.Load("c.bmp")
	sprH.Load("h.bmp")
	eng.OnKeyUp(keyUp)
}

func Show() {
	start.Show()
	levelEditor.Show()
	exit.Show()
	inv1.Show()
	inv2.Show()
	nextState = state.MAIN_MENU
	shown = true
}

func Hide() {
	start.Hide()
	levelEditor.Hide()
	exit.Hide()
	inv1.Hide()
	inv2.Hide()
	shown = false
}

func Render() state.State {
	start.Pos.Move(eng.Width()/2, 8)
	eng.SetColor(common.Color1)
	eng.Clear()
	if c {
		sprC.Render(eng.Width()/2-sprC.Width*4, 400, 8)
	}
	if h {
		sprH.Render(eng.Width()/2-sprC.Width*4+172, 365, 4)
	}
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
