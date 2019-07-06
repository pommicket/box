package eng

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// The left mouse button
	MOUSE_LEFT = iota
	// The middle mouse button
	MOUSE_MIDDLE
	// The right mouse button
	MOUSE_RIGHT
	// Another mouse button
	MOUSE_OTHER
)

var (
	// Functions to call when rendering
	onRender []func()
	// Functions to call when the window is closed
	onClose []func() bool
	// Functions to call when the user presses the mouse
	onMouseDown []func(button int, x int, y int)
	// Functions to call when the user releases the mouse
	onMouseUp []func(button int, x int, y int)
	// Functions to call when the user moves the mouse
	onMouseMove []func(x int, y int)
	// Functions to call when the user presses a key
	onKeyDown []func(keycode int)
	// Functions to call when the user releases a key
	onKeyUp []func(keycode int)
)

// Converts SDL mouse button to eng mouse button
func fromSdlMouseButton(button uint8) int {
	switch button {
	case sdl.BUTTON_LEFT:
		return MOUSE_LEFT
	case sdl.BUTTON_MIDDLE:
		return MOUSE_MIDDLE
	case sdl.BUTTON_RIGHT:
		return MOUSE_RIGHT
	}
	return MOUSE_OTHER
}

// Converts eng mouse button to SDL mouse button
func toSdlMouseButton(button int) uint8 {
	switch button {
	case MOUSE_LEFT:
		return sdl.BUTTON_LEFT
	case MOUSE_MIDDLE:
		return sdl.BUTTON_MIDDLE
	case MOUSE_RIGHT:
		return sdl.BUTTON_RIGHT
	}
	return sdl.BUTTON_LEFT
}

// Have a function be called when the window is rendered.
func OnRender(callback func()) {
	onRender = append(onRender, callback)
}

// Have a function be called when the window is closed. Iff any close function
// returns true, the window will be kept open.
func OnClose(callback func() bool) {
	onClose = append(onClose, callback)
}

// Have a function be called when the user presses a mouse button. button will
// be MOUSE_LEFT, MOUSE_MIDDLE, or MOUSE_RIGHT.
func OnMouseDown(callback func(button int, x int, y int)) {
	onMouseDown = append(onMouseDown, callback)
}

// Have a function be called when the user releases a mouse button. button will
// be MOUSE_LEFT, MOUSE_MIDDLE, or MOUSE_RIGHT.
func OnMouseUp(callback func(button int, x int, y int)) {
	onMouseUp = append(onMouseUp, callback)
}

// Have a function be called when the user moves the mouse.
func OnMouseMove(callback func(x int, y int)) {
	onMouseMove = append(onMouseMove, callback)
}

// Have a function be called when the user presses a key. key will be one of
// the constants defined in keys.go.
func OnKeyDown(callback func(key int)) {
	onKeyDown = append(onKeyDown, callback)
}

// Have a function be called when the user releases a key. key will be one of
// the constants defined in keys.go.
func OnKeyUp(callback func(key int)) {
	onKeyUp = append(onKeyUp, callback)
}

// Returns true iff this mouse button is down.
func IsMouseDown(button int) bool {
	_, _, state := sdl.GetMouseState()
	return (int(state) & (1 << (toSdlMouseButton(button) - 1))) != 0
}

// Returns the position of the mouse
func MousePos() (int, int) {
	x, y, _ := sdl.GetMouseState()
	return int(x), int(y)
}

// Returns true iff this key is down.
func IsKeyDown(key int) bool {
	state := sdl.GetKeyboardState()
	keycode := toSdlKeycode(key)
	scancode := sdl.GetScancodeFromKey(keycode)
	return state[scancode] != 0
}

// Returns true iff either control key is down
func IsCtrl() bool {
	return IsKeyDown(KEY_LCTRL) || IsKeyDown(KEY_RCTRL)
}

// Returns true iff either shift key is down
func IsShift() bool {
	return IsKeyDown(KEY_LSHIFT) || IsKeyDown(KEY_RSHIFT)
}

// Returns true iff either alt key is down
func IsAlt() bool {
	return IsKeyDown(KEY_LALT) || IsKeyDown(KEY_RALT)
}

// Handles an SDL event
func handleEvent(event sdl.Event) {
	switch event.(type) {
	case *sdl.QuitEvent:
		quit = true
		for _, f := range onClose {
			if f() {
				quit = false
			}
		}
	case *sdl.MouseButtonEvent:
		mouse := event.(*sdl.MouseButtonEvent)
		var callbacks []func(button int, x int, y int)
		if event.GetType() == sdl.MOUSEBUTTONUP {
			callbacks = onMouseUp
		} else {
			callbacks = onMouseDown
		}
		for _, f := range callbacks {
			f(fromSdlMouseButton(mouse.Button), int(mouse.X), int(mouse.Y))
		}
	case *sdl.MouseMotionEvent:
		mouse := event.(*sdl.MouseMotionEvent)
		for _, f := range onMouseMove {
			f(int(mouse.X), int(mouse.Y))
		}
	case *sdl.KeyboardEvent:
		key := event.(*sdl.KeyboardEvent)
		var callbacks []func(key int)
		if event.GetType() == sdl.KEYUP {
			callbacks = onKeyUp
		} else {
			callbacks = onKeyDown
		}
		for _, f := range callbacks {
			f(fromSdlKeycode(key.Keysym.Sym))
		}
	}
}
