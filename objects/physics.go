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

func update(dt, x, y, velX, velY float64, collidesWith map[ObjectKind]bool) (float64, float64, float64, float64, Event) {
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

	tileX := int(x)
	tileY := int(y)

	if velY != 0 {
		yVal := tileY + sgn(velY)
		inFrontOfY = []*Object{At(int(math.Floor(x)), yVal), At(int(math.Ceil(x)), yVal)}
		if inFrontOfY[0] == inFrontOfY[1] {
			inFrontOfY = inFrontOfY[:1]
		}
	}

	for _, obj := range inFrontOfY {
		obj.mutex.Lock()
		switch obj.Kind {
		case SPIKE:
			event = SPIKE_HIT
		}
		if collidesWith[obj.Kind] {
			velY = 0
			y = float64(int(y)) // Truncate
			if obj.Kind == CONVEYOR_LEFT {
				// Move to the left
				velX = -conveyor_speed
			} else if obj.Kind == CONVEYOR_RIGHT {
				// Move to the right
				velX = +conveyor_speed
			}
		}
		obj.mutex.Unlock()
	}

	if velX != 0 {
		xVal := tileX + sgn(velX)
		inFrontOfX = []*Object{At(xVal, int(math.Floor(y))), At(xVal, int(math.Ceil(y)))}
		if inFrontOfX[0] == inFrontOfX[1] {
			inFrontOfX = inFrontOfX[:1]
		}
	}

	for _, obj := range inFrontOfX {
		obj.mutex.Lock()
		switch obj.Kind {
		case SPIKE:
			event = SPIKE_HIT
		}
		if collidesWith[obj.Kind] {
			velX = 0
			x = float64(int(x)) // Truncate
		}
		obj.mutex.Unlock()
	}

	velY += gravity * dt
	return x, y, velX, velY, event
}
