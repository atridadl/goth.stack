# GOTH Stack

A modern web application stack combining:
- Go (Backend)
- HTMX (Frontend Interactivity)
- Tailwind CSS (Styling)

## Features
- Server-Side Rendering
- Real-time capabilities with SSE
- Built-in tools and utilities
- API documentation with Swagger
- Modern responsive design
- Zero client-side JavaScript (optional)

## Getting Started
1. Clone this repository
2. Copy `.env.example` to `.env`
3. Run `go mod download`
4. Start the server: `go run main.go`

## API Documentation
1. Install swag by running ```go install github.com/swaggo/swag/cmd/swag@latest```
2. Visit `/api/swagger/index.html` after starting the server

## Example Posts
Check out the `/posts` route for blogging examples

## License
MIT
