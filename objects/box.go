package objects

import (
	"github.com/pommicket/box/eng"
	"sync"
)

var boxTileX, boxTileY int
var boxX, boxY, boxLastX, boxLastY, boxVelX, boxVelY float64
var boxFrown, boxSmile eng.Sprite
var boxLock sync.Mutex

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
}

type Event int

const (
	NOTHING = Event(iota)
	ENEMY_HIT
	GOAL_REACHED
)

// Returns true if box hit spike
func updateBox(dt float64) Event {
	boxLock.Lock()
	defer boxLock.Unlock()
	var event Event
	boxLastX, boxLastY = boxX, boxY
	boxX, boxY, boxVelX, boxVelY, event = update(dt, boxX, boxY, boxVelX, boxVelY, 1, boxCollidesWith)
	if At(int(boxX), int(boxY)).GetKind() == GOAL_FLAG {
		return GOAL_REACHED
	}
	if event != NOTHING {
		return event
	}
	if anyEnemyCollidesWith(boxX, boxY) {
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
