/*
This file is part of Box.

Box is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Box is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Box.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/game"
	"github.com/pommicket/box/leveleditor"
	"github.com/pommicket/box/levelsel"
	"github.com/pommicket/box/mainmenu"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
	"github.com/veandco/go-sdl2/mix"
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

var fullscreen bool

func keyUp(key int) {
	switch key {
	case eng.KEY_f:
		fullscreen = !fullscreen
		eng.SetFullscreen(fullscreen)
		if !fullscreen {
			eng.SetSize(1920/2, 1080/2)
		}
	}
}

var music *mix.Music

func startMusic() bool {
	// Music isn't necessary, so this function won't exit if there's an error.
	if err := mix.Init(mix.INIT_OGG); err != nil {
		fmt.Println("Error initializing mixer:", err)
		return false
	}

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 1, 1024); err != nil {
		fmt.Println("Error opening audio device:", err)
		return true
	}
	var err error
	fmt.Println("Loading music...")
	music, err = mix.LoadMUS("audio/music.ogg")
	if err != nil {
		fmt.Println("Error loading music:", err)
		return true
	}

	if err := music.FadeIn(-1, 5000); err != nil {
		fmt.Println("Error playing music:", err)
		return true
	}
	return true
}

func stopMusic() {
	if music != nil {
		music.Free()
	}
	mix.HaltMusic()
	mix.Quit()
}

func main() {
	eng.PanicOnError = true
	started := startMusic()
	eng.Create("box", 1920/2, 1080/2)
	eng.OnRender(render)
	eng.OnKeyUp(keyUp)
	eng.SetSpriteDir("sprites")
	loadAll()
	switchToNewState(state.MAIN_MENU)

	go update()
	eng.Run()
	if started {
		stopMusic()
	}
}
