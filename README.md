# GOTH Stack
Go + Templates + HTMX

## Stack:
- Backend: Golang + BunRouter
- Rendering: Golang templates
- Style: TailwindCSS + DaisyUI
- Content format: Markdown

## Requirements:
- Bun (only to build styles)
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