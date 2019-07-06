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
