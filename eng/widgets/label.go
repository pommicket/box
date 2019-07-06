package widgets

import (
	"github.com/pommicket/box/eng"
)

type Label struct {
	// The position of the label.
	Pos Position
	// The pre-rendered text for the label
	text *eng.Text
	// Color for the text
	Color eng.Color
	// Is this label currently being shown?
	Shown bool
	// Has the callback been set for this label?
	callbackSet bool
}

// Sets the text of a label
func (l *Label) SetText(text string, size int) {
	if l.text == nil {
		l.text = new(eng.Text)
	} else {
		l.text.Close()
	}
	l.text.Get(text, size)
	l.Pos.W, l.Pos.H = l.text.Width, l.text.Height
}

// Shows the label. Must be called in order for the label to be rendered.
func (l *Label) Show() {
	l.Shown = true
	if !l.callbackSet {
		eng.OnRender(l.render)
		l.callbackSet = true
	}
}

// Hides the label.
func (l *Label) Hide() {
	l.Shown = false
}

// Render the label.
func (l *Label) render() {
	if !l.Shown {
		return
	}
	eng.SetColor(l.Color)
	l.text.Render(l.Pos.GetX(), l.Pos.GetY())
}

// Frees any memory allocated by this label.
func (l *Label) Close() {
	l.text.Close()
}
