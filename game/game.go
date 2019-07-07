package game

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/levels"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
	"sync"
)

var Level string
var hitSpike, goalReached bool
var nextState state.State
var mutex sync.Mutex
var eventTime float64

func Load() {

}

func Show() {
	hitSpike = false
	goalReached = false
	nextState = state.GAME
	levels.Load(Level)
}

func Hide() {

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
				levels.Load(Level)
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
