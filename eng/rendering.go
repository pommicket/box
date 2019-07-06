package eng

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"math"
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

// Draws a line from (x1, y1) to (x2, y2)
func Line(x1, y1, x2, y2 int) {
	gfx.AALineColor(sdlRenderer, int32(x1), int32(y1), int32(x2), int32(y2), color)
}

// Draws a line from (x1, y1) to (x2, y2), with the given thickness.
func ThickLine(x1, y1, x2, y2, thickness int) {
	gfx.ThickLineColor(sdlRenderer, int32(x1), int32(y1), int32(x2), int32(y2), int32(thickness), color)
}

// Draws a rectangle at (x, y) with width w and height h.
func Rectangle(x, y, w, h int, draw DrawType) {
	if draw == FILL {
		gfx.BoxColor(sdlRenderer, int32(x), int32(y), int32(x+w), int32(y+h), color)
	} else {
		gfx.RectangleColor(sdlRenderer, int32(x), int32(y), int32(x+w), int32(y+h), color)
	}
}

// Draws an ellipse centered at (x, y) with a horizontal radius of rx, and a
// vertical radius of ry.
func Ellipse(x, y, rx, ry int, draw DrawType) {
	if draw == FILL {
		gfx.FilledEllipseColor(sdlRenderer, int32(x), int32(y), int32(rx), int32(ry), color)
	} else {
		gfx.AAEllipseColor(sdlRenderer, int32(x), int32(y), int32(rx), int32(ry), color) // Anti-aliasing
	}
}

// Draws a circle centered at (x, y) with radius r.
func Circle(x, y, r int, draw DrawType) {
	Ellipse(x, y, r, r, draw)
}

// Draws a pie at (x, y), with radius r from angle start to angle end (radians).
// Currently only precise up to 1 degree.
func Pie(x, y, r int, start, end float64, draw DrawType) {
	startDegrees := int32(math.Round(start * 180 / math.Pi))
	endDegrees := int32(math.Round(end * 180 / math.Pi))
	if draw == FILL {
		gfx.FilledPieColor(sdlRenderer, int32(x), int32(y), int32(r), startDegrees, endDegrees, color)
	} else {
		gfx.PieColor(sdlRenderer, int32(x), int32(y), int32(r), startDegrees, endDegrees, color)
	}
}
