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

const (
	TOP_LEFT = iota
	TOP_MIDDLE
	TOP_RIGHT
	MIDDLE_LEFT
	MIDDLE
	MIDDLE_RIGHT
	BOTTOM_LEFT
	BOTTOM_MIDDLE
	BOTTOM_RIGHT
)

// Returns 0 if align is *_LEFT, w / 2 if align is *_MIDDLE, w if align is *_RIGHT
func alignX(w int, align int) int {
	switch align % 3 {
	case 0:
		return 0
	case 1:
		return w / 2
	case 2:
		return w
	}
	panic("Invalid align.")
	return 0
}

// Returns 0 if align is TOP_*, h / 2 if align is MIDDLE_*, h if align is BOTTOM_*
func alignY(h int, align int) int {
	switch align / 3 {
	case 0:
		return 0
	case 1:
		return h / 2
	case 2:
		return h
	}
	return 0
}

// Aligns (0, 0) to the corner given by align of an object with dimensions wxh.
func alignTo(w int, h int, align int) (int, int) {
	return alignX(w, align), alignY(h, align)
}
