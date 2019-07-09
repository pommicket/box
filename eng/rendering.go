package eng

import (
	"github.com/veandco/go-sdl2/sdl"
)

// DrawType is either DRAW (false) or FILL (true). It determines whether
// various graphics functions will draw the outline of their shapes, or fill
// them in.
type DrawType bool

const (
	DRAW = false
	FILL = true
)

var (
	// The draw color
	color sdl.Color
)

// Renders the window.
func render() {
	renderer := sdlRenderer
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	// Run render callbacks
	i := 0
	for _, f := range onRender {
		if f != nil {
			onRender[i] = f
			i++
			f()
		}
	}
	onRender = onRender[:i]
	renderer.Present()
}

// Sets the color for drawing, given red, green, blue and alpha values.
func SetRGBA(r, g, b, a uint8) {
	SetColor(Color{r, g, b, a})
}

// Sets the color for drawing, given red, green, and blue values.
// Alpha (transparency) will be set to 255.
func SetRGB(r, g, b uint8) {
	SetRGBA(r, g, b, 255)
}

// Sets the color for drawing
func SetColor(c Color) {
	color = c.toSDL()
	sdlRenderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// Clears the screen using the current draw color.
func Clear() {
	sdlRenderer.Clear()
}

// Draws a rectangle at (x, y) with width w and height h.
func Rectangle(x, y, w, h int, draw DrawType) {
	rect := sdl.Rect{int32(x), int32(y), int32(w), int32(h)}
	if draw == FILL {
		sdlRenderer.FillRect(&rect)
	} else {
		sdlRenderer.DrawRect(&rect)
	}
}
