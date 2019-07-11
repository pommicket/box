/*
This file is part of Box.

Box is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Box is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Box.  If not, see <https://www.gnu.org/licenses/>.
*/

package eng

import (
	"github.com/veandco/go-sdl2/sdl"
	"path/filepath"
)

// Stores information about a previously-loaded texture.
type textureInfo struct {
	// The actual SDL texture
	texture *sdl.Texture
	// The number of references to this texture (for keeping track of when to
	// destroy it)
	nRefs int
	// The width of the texture
	width int
	// The height of the texture
	height int
}

// A map from loaded sprite filenames to their textures.
// This allows loading multiple sprites from the same file, without having to
// re-load that file each time.
var sprites = make(map[string]textureInfo)

// The type for a sprite
type Sprite struct {
	// The SDL texture for this sprite
	texture *sdl.Texture
	// The name of the file
	filename string
	// The width of the sprite
	Width int
	// The height of the sprite
	Height int
}

// The directory where sprites are located.
var spriteDir string

// Sets the directory where sprites are located.
func SetSpriteDir(dir string) {
	spriteDir = dir
}

// Loads a sprite from an image file. The file must be a .bmp file.
func (s *Sprite) Load(filename string) error {
	filename = filepath.Join(spriteDir, filename)
	if info, ok := sprites[filename]; ok {
		// We've already loaded this texture
		info.nRefs++
		sprites[filename] = info
		s.Width = info.width
		s.Height = info.height
		s.texture = info.texture
		s.filename = filename
	} else {
		// Need to load it
		surf, err := sdl.LoadBMP(filename)
		if err != nil {
			return throwSDL("load image", "Couldn't load image.")
		}
		defer surf.Free()
		s.texture, err = sdlRenderer.CreateTextureFromSurface(surf)
		if err != nil {
			return throwSDL("create texture", "Could not create a texture for the image's surface.")
		}
		s.Width = int(surf.W)
		s.Height = int(surf.H)
		s.filename = filename
		sprites[filename] = textureInfo{s.texture, 1, s.Width, s.Height}
	}
	return nil
}

// Has this sprite been loaded?
func (s *Sprite) Loaded() bool {
	return s.texture != nil
}

// Should the next call to Sprite.Render be tinted?
var colorSprite bool

// Calling this makes the next call to Sprite.Render use the draw color to tint
// the sprite.
func ColorSprite() {
	colorSprite = true
}

var clipW, clipH int32 = -1, -1

func ClipBottomRight(w, h int) {
	clipW, clipH = int32(w), int32(h)
}

// Renders the sprite to (x, y) scaled up by a factor of scale.
func (s *Sprite) Render(x, y int, scale float64) error {
	if s.texture == nil {
		return throw("no texture", "A texture has not been created for the sprite.")
	}
	var srcrect sdl.Rect
	var width, height int
	width, height = int(float64(s.Width)*scale), int(float64(s.Height)*scale)
	srcrect.W = int32(width)
	srcrect.H = int32(height)
	dstrect := sdl.Rect{int32(x), int32(y), int32(width), int32(height)}
	if clipW >= 0 {
		srcrect.W = int32(float64(clipW) / scale)
		dstrect.W = clipW
		clipW = -1
	}
	if clipH >= 0 {
		srcrect.H = int32(float64(clipH) / scale)
		dstrect.H = clipH
		clipH = -1
	}
	if colorSprite {
		s.texture.SetColorMod(color.R, color.G, color.B)
		s.texture.SetAlphaMod(color.A)
		s.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
		err := sdlRenderer.Copy(s.texture, &srcrect, &dstrect)
		s.texture.SetColorMod(255, 255, 255)
		s.texture.SetAlphaMod(255)
		colorSprite = false
		return err
	}
	return sdlRenderer.Copy(s.texture, &srcrect, &dstrect)
}

// Closes a sprite, freeing any memory it allocated. The Width and Height fields
// are left unchanged.
func (s *Sprite) Close() {
	nRefs := sprites[s.filename].nRefs
	nRefs--
	if nRefs == 0 {
		delete(sprites, s.filename)
		s.texture.Destroy()
	} else {
		info := sprites[s.filename]
		info.nRefs--
		sprites[s.filename] = info
	}
}
