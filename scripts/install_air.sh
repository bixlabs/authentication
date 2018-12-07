#!/usr/bin/env bash

# OS=darwin for Mac
# OS=windows for Windows

OS=linux
curl -fLo ~/.air \
           https://raw.githubusercontent.com/cosmtrek/air/master/bin/${OS}/air

chmod +x ~/.air
