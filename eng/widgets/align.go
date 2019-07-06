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
