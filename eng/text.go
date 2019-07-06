package eng

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type font struct {
	// fonts[i] -> The font with size i
	fonts map[int]*ttf.Font
	// The path to the font.
	fontPath string
}

// A slice of every font which has been registered
var allFonts []font

// A type for keeping track of font IDs.
type FontId int

// Which font is currently being used
var currFont FontId

// Loads the font with the given size into f.fonts[size], and returns it.
// If f.fonts[size] exists, no action occurs.
func (f *font) loadSize(size int) (*ttf.Font, error) {
	if _, ok := f.fonts[size]; ok {
		return f.fonts[size], nil
	}
	var err error
	f.fonts[size], err = ttf.OpenFont(f.fontPath, size)
	if err != nil {
		return nil, err
	}
	return f.fonts[size], nil
}

// Loads a font for use. Returns an integer which can be used with SetFont.
// fontFilename should be a path in the resource directory
func LoadFont(fontFilename string) (FontId, error) {
	var newFont font
	newFont.fonts = make(map[int]*ttf.Font)
	newFont.fontPath = fontFilename

	// Only open size 12 font now (to make sure font path actually exists).
	ttfFont, err := ttf.OpenFont(fontFilename, 12)
	if err != nil {
		return FontId(-1), throwTTF("open font", fmt.Sprintf("Failed to open font: %v", fontFilename))
	}
	newFont.fonts[12] = ttfFont

	allFonts = append(allFonts, newFont)
	return FontId(len(allFonts) - 1), nil
}

// Sets the current font. If the font is never set, the first registered font
// will be used.
func SetFont(id FontId) {
	currFont = id
}

// The type for pre-rendered text. Pre-rendered text is useful if you want to
// either get the width and height of text before rendering it, or render the
// same text for many frames without having to re-render it each time.
type Text struct {
	// The internal SDL texture
	texture *sdl.Texture
	// The width of the text
	Width int
	// The height of the text
	Height int
}

// Gets some text and puts it into t. The color of the text at rendering will
// be used, not the color when getting. This must be called after Create.
func (t *Text) Get(text string, size int) error {
	if currFont < 0 || int(currFont) >= len(allFonts) {
		return throw("no such font", "The current font is invalid (have you loaded one yet?).")
	}
	if size < 0 {
		return throw("invalid size", "Font size must be positive.")
	}
	if size == 0 || text == "" {
		// Just don't render size/length 0 text.
		return nil
	}

	font, err := allFonts[currFont].loadSize(size)
	if err != nil {
		return throwTTF("open font", fmt.Sprint("Could not open font with size ", size))
	}
	surface, err := font.RenderUTF8Blended(text, sdl.Color{255, 255, 255, 255})
	if err != nil {
		return throwTTF("render text", fmt.Sprintf("The text %v failed to render", text))
	}
	t.Width, t.Height = int(surface.W), int(surface.H)
	defer surface.Free()
	t.texture, err = sdlRenderer.CreateTextureFromSurface(surface)
	if err != nil {
		return throwTTF("create texture", "Could not create a texture from the surface for the text.")
	}
	return nil
}

// Renders text to (x, y), with the current draw color.
func (t *Text) Render(x, y int) error {
	t.texture.SetColorMod(color.R, color.G, color.B)
	t.texture.SetAlphaMod(color.A)
	srcrect := sdl.Rect{0, 0, int32(t.Width), int32(t.Height)}
	dstrect := sdl.Rect{int32(x), int32(y), int32(t.Width), int32(t.Height)}
	return sdlRenderer.Copy(t.texture, &srcrect, &dstrect)
}

// Closes the text, freeing any memory it allocated. The Width and Height fields
// of the text are left unchanged.
func (t *Text) Close() {
	t.texture.Destroy()
}

// Draws text at (x, y) with the given size and current draw color.
// Returns the width and height of the text (as well as an error).
func RenderText(text string, x int, y int, size int) (int, int, error) {
	var t Text
	if err := t.Get(text, size); err != nil {
		return 0, 0, err
	}
	defer t.Close()
	if err := t.Render(x, y); err != nil {
		return 0, 0, err
	}
	return t.Width, t.Height, nil
}

// Close every registered font.
func closeAllFonts() {
	for _, font := range allFonts {
		for _, ttfFont := range font.fonts {
			if ttfFont != nil {
				ttfFont.Close()
			}
		}
	}
}
