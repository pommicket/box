package levels

import (
	"encoding/binary"
	"fmt"
	"github.com/pommicket/box/objects"
	"os"
)

type Version struct {
	Major int32
	Minor int32
}

var version = Version{0, 0}

// Saves current objects as a level.
func Save(filename string) {
	file, err := os.Create("game_levels/" + filename + ".box")
	if err != nil {
		fmt.Println("Error saving level as", filename, ":", err)
		os.Exit(-1)
	}

	write := func(x interface{}) {
		err := binary.Write(file, binary.LittleEndian, x)
		if err != nil {
			file.Close()
			fmt.Println("Error saving level as", filename, ":", err)
			os.Exit(-1)
		}
	}

	write(version)
	for y := 0; y < objects.TilesY; y++ {
		for x := 0; x < objects.TilesX; x++ {
			write(int32(objects.At(x, y).Kind))
		}
	}
	file.Close()
}

func Load(filename string) {
	file, err := os.Open("game_levels/" + filename + ".box")

	if err != nil {
		fmt.Println("Error loading level", filename, ":", err)
		os.Exit(-1)
	}

	read := func(x interface{}) {
		binary.Read(file, binary.LittleEndian, x)
	}

	objects.ClearAll()
	read(&version)
	for y := 0; y < objects.TilesY; y++ {
		for x := 0; x < objects.TilesX; x++ {
			var kind int32
			read(&kind)
			objects.At(x, y).Set(objects.ObjectKind(kind), true)

		}
	}
}
