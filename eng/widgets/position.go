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

package widgets

// Positions store a point in 2D space, as well as (optionally) a size.
// Positions can have parents which they are relative to.
type Position struct {
	// The x position/offset
	X int
	// The y position/offset
	Y int
	// The width of the object
	W int
	// The height of the object
	H int
	// Its parent position, to which X, Y will be added (nil for no parent).
	Parent *Position
	// Which corner should be placed at (X, Y)?
	Align int
	// Which corner of the parent should X and Y be relative to?
	ParentAlign int
}

// Sets the x and y offsets of p.
func (p *Position) Move(x int, y int) {
	p.X, p.Y = x, y
}

// Sets the parent and alignment of p.
func (p *Position) SetParent(parent *Position, align int, parentAlign int) {
	p.Parent, p.Align, p.ParentAlign = parent, align, parentAlign
}

// Gets the absolute x coordinate of the position.
func (p *Position) GetX() int {
	x := p.X
	x -= alignX(p.W, p.Align)
	if p.Parent != nil {
		x += p.Parent.GetX()
		x += alignX(p.Parent.W, p.ParentAlign)
	}
	return x
}

// Gets the absolute y coordinate of the position
func (p *Position) GetY() int {
	y := p.Y
	y -= alignY(p.H, p.Align)
	if p.Parent != nil {
		y += p.Parent.GetY()
		y += alignY(p.Parent.H, p.ParentAlign)
	}
	return y
}

// Does this position (taking into consideration width and height) include this
// point?
func (p *Position) Contains(x int, y int) bool {
	pX, pY := p.GetX(), p.GetY()
	return x >= pX && x <= pX+p.W && y >= pY && y <= pY+p.H
}
