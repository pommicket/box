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

package levels

import (
	"encoding/binary"
	"github.com/pommicket/box/objects"
	"os"
	"sync"
)

type Version struct {
	Major int32
	Minor int32
}

var version = Version{0, 0}

var loaded bool
var mutex sync.Mutex

func IsLevelLoaded() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return loaded
}

func SetLevelLoaded(l bool) {
	mutex.Lock()
	defer mutex.Unlock()
	loaded = l
}

// Saves current objects as a level.
func Save(filename string) error {
	file, err := os.Create("game_levels/" + filename + ".box")
	if err != nil {
		return err
	}

	write := func(x interface{}) error {
		err := binary.Write(file, binary.LittleEndian, x)
		if err != nil {
			file.Close()
			return err
		}
		return nil
	}

	if err := write(version); err != nil {
		return err
	}
	for y := 0; y < objects.TilesY; y++ {
		for x := 0; x < objects.TilesX; x++ {
			if err := write(int32(objects.At(x, y).Kind)); err != nil {
				return err
			}
		}
	}
	return file.Close()
}

func Load(filename string) error {
	file, err := os.Open("game_levels/" + filename + ".box")

	if err != nil {
		return err
	}

	read := func(x interface{}) error {
		return binary.Read(file, binary.LittleEndian, x)
	}

	objects.ClearAll()
	if err := read(&version); err != nil {
		return err
	}
	for y := 0; y < objects.TilesY; y++ {
		for x := 0; x < objects.TilesX; x++ {
			var kind int32
			if err := read(&kind); err != nil {
				return err
			}
			objects.At(x, y).Set(objects.ObjectKind(kind), true)

		}
	}
	SetLevelLoaded(true)
	return nil
}
