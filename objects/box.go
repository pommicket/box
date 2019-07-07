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
	PORTAL:         false,
	GOAL:           false,
	GOAL_FLAG:      false,
}

func resetBox() {
	// Resets box velocity
	boxVelX = 0
	boxVelY = 0
	boxIsStrong = false
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

// Returns true if box hit spike
func updateBox(dt float64) Event {
	boxLock.Lock()
	defer boxLock.Unlock()
	var event Event
	boxLastX, boxLastY = boxX, boxY
	boxX, boxY, boxVelX, boxVelY, event = update(dt, boxX, boxY, boxLastX, boxLastY, boxVelX, boxVelY, false, boxCollidesWith)
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
		if boxGainedStrengthTime >= 5 {
			// Powerup lasts for 5 seconds
			boxIsStrong = false
		}
	}
	if enemiesPaused {
		boxGainedPauseTime += dt
		if boxGainedPauseTime >= 5 {
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
