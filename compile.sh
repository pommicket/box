#!/bin/bash
# Build script for compiling on Linux, and cross-compiling from Linux to Windows
# If you only want to build for linux, just use:
# sudo apt install libsdl2-dev libsdl2-mixer-dev
# go get github.com/pommicket/box
# go get github.com/veandco/go-sdl2
# cd $GOPATH/src/github.com/pommicket/box
# go run main.go
# or go build -o out
# For cross-compiling, you will need MinGW:
# (Ubuntu/Debian) sudo apt install mingw-w64
# The rest of the dependencies are described here: https://github.com/veandco/go-sdl2/
# But download the 32-bit versions, rather than the 64-bit versions of SDL, etc.

# Note: This directory will be removed, then created.
# Do NOT set OUT_DIR it to something like '~'.
OUT_DIR="$HOME/Apps/box"
LIN=$OUT_DIR/linux/box
WIN=$OUT_DIR/windows/box

echo "Outputting to $OUT_DIR."


rm -rv $OUT_DIR
mkdir -p $WIN $LIN
cp ~/Apps/SDL2-2.0.9/SDL2.dll $WIN/ # Change the path to the DLL, if you want
for f in SDL2_mixer.dll libogg-0.dll libvorbis-0.dll libmpg123-0.dll libvorbisfile-3.dll libmodplug-1.dll libopusfile-0.dll libFLAC-8.dll libopus-0.dll; do
	cp ~/Apps/SDL2_mixer-2.0.4/i686-w64-mingw32/bin/$f $WIN/
done

cp -r game_levels $LIN/
cp -r game_levels $WIN/
cp -r audio $LIN/
cp -r audio $WIN/
if [ -e $LIN/game_levels/completed.txt ]; then
	rm $LIN/game_levels/completed.txt
fi
if [ -e $WIN/game_levels/completed.txt ]; then
	rm $WIN/game_levels/completed.txt
fi
cp -r sprites $LIN/
cp -r sprites $WIN/
go build -ldflags "-s -w" -o $LIN/box || exit 1
env CGO_ENABLED="1" CC="/usr/bin/i686-w64-mingw32-gcc" GOOS="windows" GOARCH="386" CGO_LDFLAGS="-lmingw32 -lSDL2 -lSDL2_mixer" CGO_CFLAGS="-D_REENTRANT" go build -ldflags "-s -w" -o $WIN/box.exe main.go
cd $WIN/..
zip -r ../box-win.zip box > /dev/null
cd -
cd $LIN/..
tar -czf ../box-linux.tar.gz box
