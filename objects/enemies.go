package objects

import (
	"sync"
)

type Enemy struct {
	mutex sync.Mutex
	x     float64
	y     float64
	velX  float64
	velY  float64
	dead  bool
}

var enemies []Enemy

var enemyCollidesWith = map[ObjectKind]bool{
	NONE:           false,
	CONVEYOR_LEFT:  true,
	CONVEYOR_RIGHT: true,
	SPIKE:          false,
	PORTAL:         false,
	GOAL:           false,
	GOAL_FLAG:      false,
}

func addEnemy(x, y int) {
	var enemy Enemy
	enemy.x = float64(x)
	enemy.y = float64(y)
	enemies = append(enemies, enemy)
}

func (e *Enemy) Update(dt float64) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.dead {
		return
	}
	e.x, e.y, e.velX, e.velY, _ = update(dt, e.x, e.y, e.velX, e.velY, -1, enemyCollidesWith)
}

func updateAllEnemies(dt float64) {
	for i := range enemies {
		enemies[i].Update(dt)
	}
}

func (e *Enemy) Render() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.dead {
		return
	}
	x, y := TilefToPixel(e.x, e.y)
	sprites[ENEMY].Render(x, y, float64(Scale()))
}

func renderAllEnemies() {
	for i := range enemies {
		enemies[i].Render()
	}
}

func (e *Enemy) CollidesWith(x, y float64) bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.dead {
		return false
	}
	overlapsX := (x >= e.x && x < e.x+1) || (e.x >= x && e.x < x+1)
	overlapsY := (y >= e.y && y < e.y+1) || (e.y >= y && e.y < y+1)
	collides := overlapsX && overlapsY
	if collides && boxLastY < e.y-1 {
		// dies
		e.dead = true
		return false
	}
	return collides
}

func anyEnemyCollidesWith(x, y float64) bool {
	for i := range enemies {
		if enemies[i].CollidesWith(x, y) {
			return true
		}
	}
	return false
}

func (e *Enemy) IsDead() bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.dead
}

func allEnemiesDead() bool {
	for i := range enemies {
		if !enemies[i].IsDead() {
			return false
		}
	}
	return true
}

func clearAllEnemies() {
	enemies = []Enemy{}
}
