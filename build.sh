#!/usr/bin/env bash

Executable_Name="terminal-exit-solver"

GOOS=windows GOARCH=amd64 go build -o ./build/${Executable_Name}-win.exe

GOOS=darwin  GOARCH=amd64 go build -o ./build/${Executable_Name}-osx
