// The main package for the game engine for AUC.
package eng

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	// The internal SDL window
	sdlWindow *sdl.Window
	// The internal SDL renderer
	sdlRenderer *sdl.Renderer
	// Will quit on next frame?
	quit bool
	// Width of window
	winWidth int
	// Height of window
	winHeight int
	// Time in milliseconds between this frame and last frame
	deltaMs uint32
)

// Create a window with the given title, width, and height.
func Create(title string, width, height int) error {
	if sdl.Init(sdl.INIT_VIDEO) != nil {
		return throwSDL("init sdl", "Failed to initialize SDL.")
	}
	if ttf.Init() != nil {
		return throwTTF("init ttf", "Failed to initialize SDL TTF.")
	}
	var err error
	sdlWindow, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		return throwSDL("create window", "Failed to create SDL Window.")
	}
	sdlRenderer, err = sdl.CreateRenderer(sdlWindow, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return throwSDL("create renderer", "Failed to create SDL renderer.")
	}
	winWidth, winHeight = width, height
	return nil
}

// Returns the width of the window
func Width() int {
	return winWidth
}

// Returns the height of the window
func Height() int {
	return winHeight
}

func SetSize(width, height int) {
	winWidth, winHeight = width, height
	sdlWindow.SetSize(int32(width), int32(height))
}

func SetFullscreen(fullscreen bool) {
	if fullscreen {
		sdlWindow.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
		var w, h int32
		w, h = sdlWindow.GetSize()
		winWidth, winHeight = int(w), int(h)
	} else {
		sdlWindow.SetFullscreen(0)
	}
}

// Returns time difference between this frame and the last one.
func DeltaTime() float64 {
	return float64(deltaMs) / 1000
}

// Opens the window. Returns when the window is closed.
func Run() {
	lastFrame := sdl.GetTicks()
	for !quit {
		for {
			event := sdl.PollEvent()
			if event == nil {
				break
			}
			handleEvent(event)
		}
		now := sdl.GetTicks()
		deltaMs = now - lastFrame
		render()
		lastFrame = now
	}
	destroy()
}

// Destroys the window.
func destroy() {
	sdlWindow.Destroy()
	sdlRenderer.Destroy()
	closeAllFonts()
	ttf.Quit()
	sdl.Quit()
}

// Sets the window to close at the end of the frame.
func Close() {
	quit = true
}
