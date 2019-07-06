package eng

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// Converts a number in the form 0xRRGGBBAA to a color.
func FromUint32(color uint32) Color {
	return Color{uint8((color >> 24) & 0xFF),
		uint8((color >> 16) & 0xFF),
		uint8((color >> 8) & 0xFF),
		uint8(color & 0xFF)}
}

// Converts an eng color to an SDL color.
func (c Color) toSDL() sdl.Color {
	return sdl.Color{c.R, c.G, c.B, c.A}
}

// Uses the red, green, and blue channels for the color.
func (c *Color) RGB(r uint8, g uint8, b uint8) {
	c.R, c.G, c.B, c.A = r, g, b, 255
}

// Uses the red, green, blue, and alpha channels for the color.
func (c *Color) RGBA(r uint8, g uint8, b uint8, a uint8) {
	c.R, c.G, c.B, c.A = r, g, b, a
}

// Uses a number in the form 0xRRGGBBAA for the color.
func (c *Color) Color(color uint32) {
	*c = FromUint32(color)
}
