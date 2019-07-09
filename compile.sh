#!/bin/bash
# Build script for compiling on Linux, and cross-compiling from Linux to Windows
# You will need MinGW, and SDL2:
# (Ubuntu/Debian) sudo apt install mingw-w64 libsdl2-dev
# The rest of the dependencies are described here: https://github.com/veandco/go-sdl2/
# But download the 32-bit versions, rather than the 64-bit versions of SDL, etc.

LIN=bin/linux/box
WIN=bin/windows/box

rm -r bin
mkdir -p $WIN $LIN
cp ~/Apps/SDL2-2.0.9/SDL2.dll $WIN/ # Change the path to the DLL, if you want
cp -r game_levels $LIN/
cp -r game_levels $WIN/
if [ -e $LIN/game_levels/completed.txt ]; then
	rm $LIN/game_levels/completed.txt
fi
if [ -e $WIN/game_levels/completed.txt ]; then
	rm $WIN/game_levels/completed.txt
fi
cp -r sprites $LIN/
cp -r sprites $WIN/
go build -o $LIN/box || exit 1
env CGO_ENABLED="1" CC="/usr/bin/i686-w64-mingw32-gcc" GOOS="windows" GOARCH="386" CGO_LDFLAGS="-lmingw32 -lSDL2" CGO_CFLAGS="-D_REENTRANT" go build -o $WIN/box.exe -x main.go
cd $WIN/..
zip -r ../box-win.zip box
cd -
cd $LIN/..
tar -czvf ../box-linux.tar.gz box
