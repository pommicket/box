package eng

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Error struct {
	// A short description of the error (e.g. "out of bounds")
	Short string
	// A longer description of the error, for printing (e.g. "The index in the array is out of bounds.")
	Long string
}

// Converts an engine error to a string.
func (e Error) Error() string {
	return e.Long
}

// Converts an engine error to a string.
func (e Error) String() string {
	return e.Long
}

// Should the engine panic when an error occurs?
var PanicOnError bool

// Panics with error message if panicOnError, otherwise returns the error.
func throw(short, long string) error {
	if PanicOnError {
		panic(long)
		return nil
	} else {
		return Error{short, long}
	}
}

// Panics with error message if panicOnError, otherwise returns the error.
// The current SDL error will be appended to the given message.
func throwSDL(short, long string) error {
	return throw(short, long+", SDL Error: "+sdl.GetError().Error())
}
