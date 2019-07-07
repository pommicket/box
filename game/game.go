package game

import (
	"fmt"
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/levels"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
	"os"
	"sync"
)

var Level string
var hitSpike, goalReached bool
var nextState state.State
var mutex sync.Mutex
var eventTime float64
var shown bool

func Load() {
	eng.OnKeyUp(keyUp)
}

func ResetLevel() {
	if err := levels.Load(Level); err != nil {
		fmt.Println("Error loading level:", err)
		os.Exit(-1)
	}
}

func Show() {
	hitSpike = false
	goalReached = false
	nextState = state.GAME
	shown = true
	ResetLevel()
}

func Hide() {
	shown = false
}

func Update(dt float64) {
	mutex.Lock()
	defer mutex.Unlock()
	if hitSpike || goalReached {
		if common.AbsTime()-eventTime >= 1 {
			if hitSpike {
				// After 1 second, reset level
				common.SetGameSpeed(1)
				hitSpike = false
				ResetLevel()
			} else {
				// Go to level select
				nextState = state.LEVEL_SELECT
			}
		}
	} else {
		switch objects.UpdateAll(dt) {
		case objects.ENEMY_HIT:
			eventTime = common.AbsTime()
			hitSpike = true
			common.PauseGame()
		case objects.GOAL_REACHED:
			eventTime = common.AbsTime()
			goalReached = true
			common.PauseGame()
		case objects.GOT_GRAVITY:
			objects.ReverseGravity()
		case objects.NOTHING:
		}
	}

}

func Render() state.State {
	eng.SetColor(common.Color1)
	eng.Clear()
	objects.RenderAll(false)
	mutex.Lock()
	defer mutex.Unlock()
	if hitSpike {
		objects.RenderBoxFrown()
	} else if goalReached {
		objects.RenderBoxSmile()
	}
	return nextState
}

var pausedUsingP bool

func keyUp(key int) {
	if !shown {
		return
	}
	switch key {
	case eng.KEY_r:
		if !common.IsPaused() {
			ResetLevel()
		}
	case eng.KEY_ESCAPE:
		mutex.Lock()
		defer mutex.Unlock()
		nextState = state.LEVEL_SELECT
	case eng.KEY_p:
		if common.IsPaused() {
			if pausedUsingP {
				common.SetGameSpeed(1)
			}
		} else {
			common.PauseGame()
			pausedUsingP = true
		}
	}
}
