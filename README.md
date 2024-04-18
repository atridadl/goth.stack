# atri.dad
This is my personal website!

## Stack:
- Backend: Golang + Echo
- Rendering: Golang templates
- Style: TailwindCSS + DaisyUI (No JS Needed)
- Content format: Markdown

## Requirements:
- Golang 1.22.0

## Instructions:
1. Run go get
2. Duplicate the .env.example file and call it .env
3. Fill out the .env values
4. Run ```go install github.com/cosmtrek/air@latest``` to download Air for live reload
5. Run ```air``` to start the dev server

_Note that on MacOS, you need to right click and open the appropriate tailwind executable before you can run StyleGen. This is a limitation of running unsigned binaries in MacOS. Blame Tim Apple._

## Tests
Without Coverage: `go test atri.dad/lib`
With Coverage: `go test atri.dad/lib -cover`
