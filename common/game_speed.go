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

package common

import (
	"sync"
	"time"
)

// dt will be multiplied by this
var gameSpeed float64 = 1
var mutex sync.RWMutex

func SetGameSpeed(speed float64) {
	mutex.Lock()
	defer mutex.Unlock()
	gameSpeed = speed
}

func PauseGame() {
	SetGameSpeed(0)
}

func GetGameSpeed() float64 {
	mutex.RLock()
	defer mutex.RUnlock()
	return gameSpeed
}

func IsPaused() bool {
	return GetGameSpeed() == 0
}

func AbsTime() float64 {
	// Absolute time, independent of game speed
	return float64(time.Now().UTC().UnixNano()) / 1e9
}
