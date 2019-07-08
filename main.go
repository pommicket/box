package main

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/game"
	"github.com/pommicket/box/leveleditor"
	"github.com/pommicket/box/levelsel"
	"github.com/pommicket/box/mainmenu"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
	"time"
)

func loadAll() {
	objects.Load()
	mainmenu.Load()
	leveleditor.Load()
	levelsel.Load()
	game.Load()
}

func switchToNewState(newState state.State) {
	if state.Get() == newState {
		return
	}
	switch state.Get() {
	case state.MAIN_MENU:
		mainmenu.Hide()
	case state.LEVEL_EDITOR:
		leveleditor.Hide()
	case state.LEVEL_SELECT:
		levelsel.Hide()
	case state.GAME:
		game.Hide()
		if game.Completed {
			levelsel.Completed(game.Level)
		}
	}
	state.Set(newState)
	switch newState {
	case state.MAIN_MENU:
		mainmenu.Show()
	case state.LEVEL_EDITOR:
		leveleditor.Show()
	case state.LEVEL_SELECT:
		levelsel.Show()
	case state.GAME:
		game.Show()
	case state.EXIT:
		eng.Close()
	}
}

func render() {
	var newState state.State
	switch state.Get() {
	case state.MAIN_MENU:
		newState = mainmenu.Render()
	case state.LEVEL_EDITOR:
		newState = leveleditor.Render()
	case state.LEVEL_SELECT:
		newState = levelsel.Render()
	case state.GAME:
		newState = game.Render()
	}
	switchToNewState(newState)
}

var updateEvery time.Duration = 10 * time.Millisecond

func update() {
	lastTime := time.Now().UTC().UnixNano()
	for {
		now := time.Now().UTC().UnixNano()
		dt := float64(now-lastTime) / 1e9
		switch state.Get() {
		case state.GAME:
			game.Update(dt * common.GetGameSpeed())
		}
		lastTime = now
		time.Sleep(updateEvery)
	}
}

func main() {
	eng.PanicOnError = true
	eng.Create("box", 1920/2, 1080/2)
	eng.OnRender(render)
	eng.SetSpriteDir("sprites")
	loadAll()
	switchToNewState(state.MAIN_MENU)
	go update()
	eng.Run()
}
