package state

import (
	"sync"
)

type State int

const (
	none = State(iota)
	MAIN_MENU
	LEVEL_EDITOR
	LEVEL_SELECT
	GAME
	EXIT
)

var current State
var mutex sync.RWMutex

func Get() State {
	mutex.RLock()
	defer mutex.RUnlock()
	return current
}

func Set(s State) {
	mutex.Lock()
	defer mutex.Unlock()
	current = s
}
