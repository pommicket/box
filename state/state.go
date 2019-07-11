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
