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
	"io"
	"os"
)

var levelNames []string
var levelsCompleted map[string]bool
var buttons []widgets.Position // Use pos for custom rendering
var digits []eng.Sprite
var shown bool
var nextState state.State
var winBtn widgets.Button

func AreAllCompleted() bool {
	for _, v := range levelsCompleted {
		if !v {
			return false
		}
	}
	return true
}

func Load() {
	levelsCompleted = make(map[string]bool)
	readLines := func(filename string) ([]string, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			lines = append(lines, line)
		}
		return lines, file.Close()
	}
	// Read level listing
	var err error
	levelNames, err = readLines("game_levels/listing.txt")
	if err != nil {
		fmt.Println("Error opening listing file:", err)
		os.Exit(-1)
	}
	for _, name := range levelNames {
		levelsCompleted[name] = false
	}

	// Read completed levels
	completed, err := readLines("game_levels/completed.txt")
	if !os.IsNotExist(err) {
		// File exists
		if err != nil {
			// There was some other error (permission denied, etc.)
			fmt.Println("Error opening completed file:", err)
			os.Exit(-1)
		}
		for _, name := range completed {
			levelsCompleted[name] = true
		}
	}

	x := 8
	y := 8
	buttons = make([]widgets.Position, 0, len(levelNames))
	for range levelNames {
		var pos widgets.Position
		pos.Move(x, y)
		pos.W = 36
		pos.H = 36
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
	winBtn.LoadAll("win.bmp")
	winBtn.OnClick = func() {
		game.Level = "i"
		nextState = state.GAME
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

func showNumberBox(pos *widgets.Position, n int, completed bool) {
	scale := objects.Scale()
	W := pos.W * scale
	H := pos.H * scale
	x := pos.GetX() * scale
	y := pos.GetY() * scale
	scale *= 3
	eng.SetColor(common.Color2)
	if completed {
		eng.Rectangle(x, y, W, H, eng.FILL)
	} else {
		eng.Rectangle(x, y, W, H, eng.DRAW)
	}
	digit1 := &digits[n/10]
	digit2 := &digits[n%10]
	y += H/2 - (digit1.Height*scale)/2
	dx := digit1.Width*scale + scale // dx between first digit and second
	digitsWidth := dx + digit2.Width*scale
	x += W/2 - digitsWidth/2
	if completed {
		eng.SetColor(common.Color1)
	} else {
		eng.SetColor(common.Color2)
	}
	eng.ColorSprite()
	digit1.Render(x, y, float64(scale))
	eng.ColorSprite()
	digit2.Render(x+dx, y, float64(scale))
}

func mouseUp(button, x, y int) {
	if !shown {
		return
	}
	if button != eng.MOUSE_LEFT {
		return
	}
	scale := objects.Scale()
	for i := range buttons {
		buttons[i].X *= scale
		buttons[i].Y *= scale
		buttons[i].W *= scale
		buttons[i].H *= scale
		if buttons[i].Contains(x, y) {
			game.Level = levelNames[i]
			nextState = state.GAME
			buttons[i].X /= scale
			buttons[i].Y /= scale
			buttons[i].W /= scale
			buttons[i].H /= scale
			return
		}
		buttons[i].X /= scale
		buttons[i].Y /= scale
		buttons[i].W /= scale
		buttons[i].H /= scale
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
	if AreAllCompleted() {
		winBtn.Show()
	} else {
		winBtn.Hide()
	}
	for i := range buttons {
		showNumberBox(&buttons[i], i, levelsCompleted[levelNames[i]])
	}
	winBtn.Render()
	winBtn.Scale = 4 * float64(objects.Scale())
	x := eng.Width()/2 - winBtn.Width()/2
	y := eng.Height() - winBtn.Height() - int(8*objects.Scale())
	winBtn.Pos.Move(x, y)
	return nextState
}

func Completed(levelName string) {
	levelsCompleted[levelName] = true
	file, err := os.Create("game_levels/completed.txt")
	if err != nil {
		fmt.Println("Error opening completed file:", err)
		os.Exit(-1)
	}
	defer file.Close()
	for name, completed := range levelsCompleted {
		if completed {
			_, err = io.WriteString(file, name+"\n")
			if err != nil {
				fmt.Println("Error writing to completed file:", err)
				os.Exit(-1)
			}
		}
	}

}
