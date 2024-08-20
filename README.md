# atri.dad
This is my personal website!

## Stack:
- Backend: Golang + Echo
- Rendering: Golang templates
- Style: TailwindCSS + DaisyUI
- Content format: Markdown

## Requirements:
- Golang 1.22.0

## Instructions:
1. Run ```go get```
2. Duplicate the .env.example file and call it .env
3. Fill out the required .env values
4. Run ```go install github.com/cosmtrek/air@latest``` to download Air for live reload
5. Run ```air``` to start the dev server (macOS and Linux only)

_Note that on MacOS, you need to right click and open the appropriate tailwind executable before you can run StyleGen. This is a limitation of running unsigned binaries in MacOS. Blame Tim Apple._
_Also note that I will not provide steps for Windows._
