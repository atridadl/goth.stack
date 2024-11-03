// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "Get a pong response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping the server",
                "responses": {
                    "200": {
                        "description": "Pong!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/copy": {
            "get": {
                "description": "Returns an SVG of a copy icon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Get copy icon SVG",
                "responses": {
                    "200": {
                        "description": "SVG content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sse": {
            "get": {
                "description": "Establishes a Server-Sent Events connection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/event-stream"
                ],
                "tags": [
                    "sse"
                ],
                "summary": "Server-Sent Events endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel name",
                        "name": "channel",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event stream",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tools/resize": {
            "post": {
                "description": "Resizes an uploaded image to specified dimensions",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "tools"
                ],
                "summary": "Resize an image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file to resize",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Target width",
                        "name": "width",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Target height",
                        "name": "height",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resized image",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tools/sendsse": {
            "post": {
                "description": "Sends a message to a specified SSE channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sse",
                    "tools"
                ],
                "summary": "Send SSE message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel name",
                        "name": "channel",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Message to send",
                        "name": "message",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "GOTH Stack API",
	Description:      "This is the API for GOTH Stack - Go + HTMX + Tailwind",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
