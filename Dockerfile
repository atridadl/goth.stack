FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install UPX
RUN apk add --no-cache upx

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-s -w -extldflags "-static"' \
    -tags netgo,osusergo \
    -o /go/bin/app

# Final stage
FROM scratch

COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]
