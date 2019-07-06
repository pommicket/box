package levelsel

import (
	"bufio"
	"fmt"
	"github.com/pommicket/box/common"
	"github.com/pommicket/box/eng"
	"github.com/pommicket/box/eng/widgets"
	"github.com/pommicket/box/game"
	"github.com/pommicket/box/objects"
	"github.com/pommicket/box/state"
	"os"
)

var levelNames []string
var buttons []widgets.Position // Use pos for custom rendering
var digits []eng.Sprite
var shown bool
var nextState state.State

func Load() {
	// Read level listing
	listingFile, err := os.Open("game_levels/listing.txt")
	if err != nil {
		fmt.Println("Error opening listing file:", err)
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(listingFile)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		levelNames = append(levelNames, line)
	}

	x := 8
	y := 8
	buttons = make([]widgets.Position, 0, len(levelNames))
	for range levelNames {
		var pos widgets.Position
		pos.Move(x, y)
		pos.W = 36 * objects.Scale()
		pos.H = 36 * objects.Scale()
		buttons = append(buttons, pos)
		x += 8 + pos.W
		if x+pos.W >= eng.Width() {
			y += pos.H + 8
			x = 8
		}
	}

	digits = make([]eng.Sprite, 10)
	for i := range digits {
		digits[i].Load(fmt.Sprintf("%v.bmp", i))
	}
	eng.OnMouseUp(mouseUp)
	eng.OnKeyUp(keyUp)
}

func Show() {
	shown = true
	nextState = state.LEVEL_SELECT
}

func Hide() {
	shown = false
}

func showNumberBox(pos *widgets.Position, n int) {
	scale := objects.Scale()
	W := pos.W
	H := pos.H
	x := pos.GetX()
	y := pos.GetY()
	scale *= 3
	eng.SetColor(common.Color2)
	eng.Rectangle(x, y, W, H, eng.DRAW)
	digit1 := &digits[n/10]
	digit2 := &digits[n%10]
	y += H/2 - (digit1.Height*scale)/2
	dx := digit1.Width*scale + scale // dx between first digit and second
	digitsWidth := dx + digit2.Width*scale
	x += W/2 - digitsWidth/2
	digit1.Render(x, y, float64(scale))
	digit2.Render(x+dx, y, float64(scale))
}

func mouseUp(button, x, y int) {
	if !shown {
		return
	}
	if button != eng.MOUSE_LEFT {
		return
	}
	for i := range buttons {
		if buttons[i].Contains(x, y) {
			game.Level = levelNames[i]
			nextState = state.GAME
			return
		}
	}
}

func keyUp(key int) {
	if !shown {
		return
	}
	if key == eng.KEY_ESCAPE {
		nextState = state.MAIN_MENU
	}
}

func Render() state.State {
	eng.SetColor(common.Color1)
	eng.Clear()
	for i := range buttons {
		showNumberBox(&buttons[i], i)
	}
	return nextState
}
