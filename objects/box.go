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

package objects

import (
	"github.com/pommicket/box/eng"
	"sync"
)

var boxTileX, boxTileY int
var boxX, boxY, boxLastX, boxLastY, boxVelX, boxVelY float64
var boxFrown, boxSmile eng.Sprite
var boxLock sync.Mutex
var boxIsStrong bool
var boxGainedStrengthTime float64
var boxGainedPauseTime float64

var boxCollidesWith = map[ObjectKind]bool{
	NONE:           false,
	CONVEYOR_LEFT:  true,
	CONVEYOR_RIGHT: true,
	SPIKE:          true,
	WALL:           true,
	PORTAL:         false,
	GOAL:           false,
	GOAL_FLAG:      false,
}

func resetBox() {
	// Resets box velocity
	boxVelX = 0
	boxVelY = 0
	boxIsStrong = false
	enemiesPaused = false
}

type Event int

const (
	NOTHING = Event(iota)
	ENEMY_HIT
	GOAL_REACHED
	GOT_GRAVITY
	GOT_STRENGTH
	GOT_PAUSE
)

const powerupTime = 8

// Returns true if box hit spike
func updateBox(dt float64) Event {
	boxLock.Lock()
	defer boxLock.Unlock()
	var event Event
	boxLastX, boxLastY = boxX, boxY
	boxX, boxY, boxVelX, boxVelY, event = update(dt, boxX, boxY, boxVelX, boxVelY, false, boxCollidesWith)
	if At(int(boxX), int(boxY)).GetKind() == GOAL_FLAG {
		return GOAL_REACHED
	}
	if event == ENEMY_HIT && boxIsStrong {
		event = NOTHING // for spikes
	}
	if event == GOT_STRENGTH {
		boxIsStrong = true
		boxGainedStrengthTime = 0
		event = NOTHING // Higher up functions don't need to know about this
	}
	if event == GOT_PAUSE {
		enemiesPaused = true
		boxGainedPauseTime = 0
		event = NOTHING
	}
	if boxIsStrong {
		boxGainedStrengthTime += dt
		if boxGainedStrengthTime >= powerupTime {
			boxIsStrong = false
		}
	}
	if enemiesPaused {
		boxGainedPauseTime += dt
		if boxGainedPauseTime >= powerupTime {
			enemiesPaused = false
		}
	}
	if event != NOTHING {
		return event
	}
	if anyEnemyCollidesWith(boxX, boxY, boxIsStrong) {
		return ENEMY_HIT
	}
	return NOTHING
}

func teleportBox(x, y int) {
	boxLock.Lock()
	defer boxLock.Unlock()
	boxX = float64(x)
	boxY = float64(y)
}

func renderBox() {
	boxLock.Lock()
	defer boxLock.Unlock()
	sprite := sprites[BOX]
	scale := float64(Scale())
	pixelX, pixelY := TilefToPixel(boxX, boxY)
	sprite.Render(pixelX, pixelY, scale)
}

func renderBoxWith(sprite *eng.Sprite) {
	scale := float64(Scale())
	pixelX, pixelY := TilefToPixel(boxX, boxY-1)
	sprite.Render(pixelX, pixelY, scale)
}

func RenderBoxFrown() {
	if !boxFrown.Loaded() {
		boxFrown.Load("frown.bmp")
	}
	renderBoxWith(&boxFrown)
}

func RenderBoxSmile() {
	if !boxSmile.Loaded() {
		boxSmile.Load("smile.bmp")
	}
	renderBoxWith(&boxSmile)
}
