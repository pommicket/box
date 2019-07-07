package objects

import (
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/state"
	"math"
	"sync"
)

type ObjectKind int

const (
	NONE = ObjectKind(iota)
	BOX  // NOTE: box and enemy are only objects in the level editor
	ENEMY
	SPIKE
	CONVEYOR_LEFT
	CONVEYOR_RIGHT
	PORTAL
	GOAL
	GOAL_FLAG
	POWERUP_GRAVITY
	POWERUP_STRENGTH
)

var spriteFilenames = map[ObjectKind]string{
	BOX:              "box.bmp",
	ENEMY:            "enemy.bmp",
	SPIKE:            "spike.bmp",
	CONVEYOR_LEFT:    "conveyor.bmp",
	CONVEYOR_RIGHT:   "conveyor.bmp",
	PORTAL:           "portal.bmp",
	GOAL:             "goal.bmp",
	GOAL_FLAG:        "goal_flag.bmp",
	POWERUP_GRAVITY:  "gravity.bmp",
	POWERUP_STRENGTH: "strength.bmp",
}
var sprites map[ObjectKind]*eng.Sprite
var leftSprite, rightSprite eng.Sprite

type Object struct {
	Kind   ObjectKind
	X, Y   int
	arrowX float64
	// Support multithreading
	mutex sync.Mutex
}

var objects [][]Object

const TilesX = 48
const TilesY = 27

func Load() {
	sprites = make(map[ObjectKind]*eng.Sprite)
	for kind, filename := range spriteFilenames {
		sprites[kind] = new(eng.Sprite)
		sprites[kind].Load(filename)
	}
	leftSprite.Load("left.bmp")
	rightSprite.Load("right.bmp")
	objects = make([][]Object, TilesY)
	for y := range objects {
		objects[y] = make([]Object, TilesX)
		for x := range objects[y] {
			objects[y][x].X = x
			objects[y][x].Y = y
		}
	}
	eng.OnMouseUp(mouseUp)
}

func TileToPixel(x int, y int) (int, int) {
	return x * TileSize(), y * TileSize()
}

func TilefToPixel(x float64, y float64) (int, int) {
	return int(x * float64(TileSize())), int(y * float64(TileSize()))
}

func PixelToTile(x int, y int) (int, int) {
	return x / TileSize(), y / TileSize()
}

func At(x int, y int) *Object {
	// uses modulo
	x = ((x % TilesX) + TilesX) % TilesX
	y = ((y % TilesY) + TilesY) % TilesY
	return &objects[y][x]
}

func Scale() int {
	return TileSize() / 20
}

func (o *Object) Set(kind ObjectKind, checkBox bool) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if state.Get() == state.LEVEL_EDITOR {
		if kind == BOX && checkBox {
			// Make sure there's only one box
			if objects[boxTileY][boxTileX].Kind == BOX {
				objects[boxTileY][boxTileX].Kind = NONE
			}
			boxTileX = o.X
			boxTileY = o.Y
		}
	} else if kind == BOX {
		// Don't actually place box
		boxX = float64(o.X)
		boxY = float64(o.Y)
		return
	} else if kind == ENEMY {
		// Don't actually place enemy
		addEnemy(o.X, o.Y)
		return
	}
	o.Kind = kind
	o.arrowX = 0.5
}

func (o *Object) GetKind() ObjectKind {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	return o.Kind
}

func TileSize() int {
	return eng.Width() / TilesX
}

func (o *Object) render(ghost bool) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.Kind == NONE {
		return
	}
	sprite := sprites[o.Kind]
	scale := float64(TileSize()) / float64(sprite.Width)
	pixelX, pixelY := TileToPixel(o.X, o.Y)
	if ghost {
		eng.SetRGBA(255, 255, 255, 127)
		eng.ColorSprite()
	}
	sprite.Render(pixelX, pixelY, scale)
	// Draw arrow
	var arrowSprite eng.Sprite

	if o.Kind == CONVEYOR_LEFT {
		arrowSprite = leftSprite
	} else if o.Kind == CONVEYOR_RIGHT {
		arrowSprite = rightSprite
	}

	if arrowSprite.Loaded() {
		x := o.X*TileSize() + int(o.arrowX*float64(TileSize()))
		y := o.Y*TileSize() + TileSize()/2
		if ghost {
			eng.SetRGBA(255, 255, 255, 127)
			eng.ColorSprite()
		}
		arrowSprite.Render(x, y, scale)
	}
}

func (o *Object) Update(dt float64) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	switch o.Kind {
	case CONVEYOR_LEFT:
		o.arrowX -= dt / 2
	case CONVEYOR_RIGHT:
		o.arrowX += dt / 2
	case GOAL:
		if allEnemiesDead() {
			o.Kind = GOAL_FLAG
		}
	}
	o.arrowX = math.Mod(o.arrowX, 1)
}

func (o *Object) RenderGhost() {
	o.render(true)
}

func (o *Object) Render() {
	o.render(false)
}

func RenderAll(showGrid bool) {
	for y := range objects {
		for x := range objects[y] {
			if showGrid {
				eng.SetRGBA(common.Color2.R, common.Color2.G, common.Color2.B, 25)
				eng.Rectangle(x*TileSize(), y*TileSize(), TileSize(), TileSize(), eng.DRAW)
			}
			At(x, y).Render()
		}
	}
	if state.Get() == state.GAME {
		renderBox()
		renderAllEnemies()
	}
}

func UpdateAll(dt float64) Event {
	for y := range objects {
		for x := range objects[y] {
			At(x, y).Update(dt)
		}
	}
	if state.Get() == state.GAME {
		updateAllEnemies(dt)
		return updateBox(dt)
	}
	return NOTHING
}

func ClearAll() {
	for y := range objects {
		for x := range objects[y] {
			objects[y][x].Kind = NONE
			objects[y][x].arrowX = 0.5
		}
	}
	resetBox()
	clearAllEnemies()
	common.SetGameSpeed(1)
}

func mouseUp(button, x, y int) {
	if state.Get() != state.GAME || button != eng.MOUSE_LEFT || common.IsPaused() {
		return
	}

	tileX, tileY := PixelToTile(x, y)
	o := At(tileX, tileY)
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if At(tileX, tileY).Kind != PORTAL {
		return
	}
	teleportBox(tileX, tileY)
}
