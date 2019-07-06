package objects

// import (
// 	"sync"
// )
//
// type Enemy struct {
// 	mutex sync.Mutex
// 	x float64
// 	y float64
// 	velX float64
// 	velY float64
// }
//
//
// var enemies []Enemy
//
// func addEnemy(x, y int) {
// 	var enemy Enemy
// 	enemy.x = x
// 	enemy.y = y
// 	enemies = append(enemies, enemy)
// }
//
// func (e *Enemy) Update(dt float64) {
// 	e.mutex.Lock()
// 	defer e.mutex.Unlock()
// 	e.x += dt * e.velX
// 	e.y += dt * e.velY + 0.5 * dt * dt * gravity
// }
