package widgets

import (
	"github.com/pommicket/box/eng"
)

type Button struct {
	// Position of the button
	Pos Position
	// Main sprite for the button
	Sprite eng.Sprite
	// Sprite to use when user is hovering over button
	SpriteHover eng.Sprite
	// Sprite to use while user is clicking button
	SpriteActive eng.Sprite
	// Scale to draw the sprite at. If 0, will use 1 instead.
	Scale float64
	// Function to be called when the button is clicked. nil for no function.
	OnClick func()
	// Is the mouse currently being hovered over the button?
	// This is not updated if the button is hidden.
	Hovering bool
	// Is the button currently being pressed?
	// This is not updated if the button is hidden.
	Active bool
	// Is this button shown?
	Shown bool
	// Have the callbacks been set for this button?
	callbacksSet bool
}

// Updates the dimensions of the button (Pos.W, Pos.H)
func (b *Button) updateDims() {
	b.Pos.W = int(float64(b.Sprite.Width) * b.Scale)
	b.Pos.H = int(float64(b.Sprite.Height) * b.Scale)
}

// Sets all sprites to the same sprite
func (b *Button) SetAll(sprite eng.Sprite) {
	b.Sprite = sprite
	b.SpriteHover = sprite
	b.SpriteActive = sprite
}

// Loads all sprites using a filename
func (b *Button) LoadAll(filename string) {
	b.Sprite.Load(filename)
	b.SpriteHover = b.Sprite
	b.SpriteActive = b.Sprite
}

// Load sprites using filenames [filename].bmp, [filename]_hover.bmp,
// [filename]_active.bmp
func (b *Button) LoadWithSuffixes(filename string) {
	b.Sprite.Load(filename + ".bmp")
	b.SpriteHover.Load(filename + "_hover.bmp")
	b.SpriteActive.Load(filename + "_active.bmp")
}

// Shows the button
func (b *Button) Show() {
	if !b.callbacksSet {
		eng.OnRender(b.render)
		eng.OnMouseMove(b.mouseMove)
		eng.OnMouseDown(b.mouseDown)
		eng.OnMouseUp(b.mouseUp)
		b.callbacksSet = true
	}
	b.Shown = true
	x, y := eng.MousePos()
	b.mouseMove(x, y) // Check for hovering, etc.
}

// Hides the button
func (b *Button) Hide() {
	b.Shown = false
}

// Handle mouse motion for the button
func (b *Button) mouseMove(x, y int) {
	if !b.Shown {
		return
	}
	b.updateDims()
	b.Hovering = b.Pos.Contains(x, y)
}

// Handle a mouse down event for the button
func (b *Button) mouseDown(button, x, y int) {
	if !b.Shown {
		return
	}
	if button != eng.MOUSE_LEFT { // Only care about left mouse clicks.
		return
	}
	if b.Hovering {
		b.Active = true
	}
}

// Handle a mouse up event for the button
func (b *Button) mouseUp(button, x, y int) {
	if !b.Shown {
		return
	}
	if button != eng.MOUSE_LEFT { // Only care about left mouse clicks.
		return
	}
	if b.Active {
		// Button was just released
		if b.OnClick != nil {
			b.OnClick()
		}
	}
	b.Active = false
}

// Render the button
func (b *Button) render() {
	if !b.Shown {
		return
	}
	if !b.Sprite.Loaded() {
		// We don't have a sprite loaded yet.
		// Don't render.
		return
	}
	scale := b.Scale
	if scale == 0 {
		scale = 1
	}
	sprite := b.Sprite
	if b.Active {
		sprite = b.SpriteActive
	} else if b.Hovering {
		sprite = b.SpriteHover
	}
	b.updateDims()
	// Render it!
	sprite.Render(b.Pos.GetX(), b.Pos.GetY(), scale)
}
