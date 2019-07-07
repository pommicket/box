package objects

import (
	"math"
)

func sgn(x float64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

const conveyor_speed = 2 // In tiles / sec
const gravity = 5

func update(dt, x, y, velX, velY, conveyorSpeedMultiplier float64, collidesWith map[ObjectKind]bool) (float64, float64, float64, float64, Event) {
	x += dt * velX
	y += dt*velY + 0.5*dt*dt*gravity

	var inFrontOfX, inFrontOfY []*Object // Objects in directions of motion of object
	event := NOTHING
	if x < 0 {
		x += TilesX
	} else if x >= TilesX {
		x -= TilesX
	}
	if y < 0 {
		y += TilesY
	} else if y >= TilesY {
		y -= TilesY
	}

	if velY != 0 {
		var yVal int
		if velY < 0 {
			yVal = int(y)
		} else {
			yVal = int(y) + 1
		}
		inFrontOfY = []*Object{At(int(math.Floor(x)), yVal), At(int(math.Ceil(x)), yVal)}
		if inFrontOfY[0] == inFrontOfY[1] {
			inFrontOfY = inFrontOfY[:1]
		}
	}

	for _, obj := range inFrontOfY {
		obj.mutex.Lock()
		switch obj.Kind {
		case SPIKE:
			event = ENEMY_HIT
		}
		if collidesWith[obj.Kind] {
			if velY > 0 {
				y = math.Floor(y)
			} else {
				y = math.Ceil(y)
			}
			velY = 0
			if obj.Kind == CONVEYOR_LEFT {
				// Move to the left
				velX = -conveyor_speed * conveyorSpeedMultiplier
			} else if obj.Kind == CONVEYOR_RIGHT {
				// Move to the right
				velX = +conveyor_speed * conveyorSpeedMultiplier
			}
		}
		obj.mutex.Unlock()
	}

	if velX != 0 {
		var xVal int
		if velX < 0 {
			xVal = int(x)
		} else {
			xVal = int(math.Ceil(x))
		}
		inFrontOfX = []*Object{At(xVal, int(math.Floor(y))), At(xVal, int(math.Ceil(y)))}
		if inFrontOfX[0] == inFrontOfX[1] {
			inFrontOfX = inFrontOfX[:1]
		}
	}

	for _, obj := range inFrontOfX {
		obj.mutex.Lock()
		switch obj.Kind {
		case SPIKE:
			event = ENEMY_HIT
		}
		if collidesWith[obj.Kind] {
			if velX > 0 {
				x = math.Floor(x)
			} else {
				x = math.Ceil(x)
			}
			velX = 0
		}
		obj.mutex.Unlock()
	}

	velY += gravity * dt
	return x, y, velX, velY, event
}