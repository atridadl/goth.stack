![goth](https://github.com/atridadl/goth.stack/assets/88056492/7c973d6a-fcf3-41fd-a119-1a81da52b342)
# GOTH Stack


Go + Templates + HTMX

## Stack:
- Backend: Golang + Echo
- Rendering: Golang templates
- Style: TailwindCSS + DaisyUI (No JS Needed)
- Content format: Markdown

## Requirements:
- Golang 1.21.6 or newer

## Instructions:
1. Run go get
2. Duplicate the .env.example file and call it .env
3. Fill out the .env values
4. Run ```go install github.com/cosmtrek/air@latest``` to download Air for live reload
5. Run ```air``` to start the dev server

## Tests
Without Coverage: `go test goth.stack/lib`
With Coverage: `go test goth.stack/lib -cover`
